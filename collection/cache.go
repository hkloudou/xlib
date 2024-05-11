package collection

import (
	"container/list"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hkloudou/xlib/xlog"
	"github.com/hkloudou/xlib/xmath"
	"github.com/hkloudou/xlib/xsync"
)

const (
	defaultCacheName = "proc"
	slots            = 300
	statInterval     = time.Minute
	// make the expiry unstable to avoid lots of cached items expire at the same time
	// make the unstable expiry to be [0.95, 1.05] * seconds
	expiryDeviation = 0.05
)

var emptyLruCache = emptyLru{}

type (
	// CacheOption defines the method to customize a Cache.
	CacheOption[T any] func(cache *Cache[T])

	// A Cache object is an in-memory cache.
	Cache[T any] struct {
		name           string
		lock           sync.Mutex
		data           map[string]T
		expire         time.Duration
		timingWheel    *TimingWheel
		lruCache       lru
		barrier        xsync.SingleFlight[T]
		unstableExpiry xmath.Unstable
		stats          *cacheStat
	}
)

// NewCache returns a Cache with given expire.
func NewCache[T any](expire time.Duration, opts ...CacheOption[T]) (*Cache[T], error) {
	cache := &Cache[T]{
		data:           make(map[string]T),
		expire:         expire,
		lruCache:       emptyLruCache,
		barrier:        xsync.NewSingleFlight[T](),
		unstableExpiry: xmath.NewUnstable(expiryDeviation),
	}

	for _, opt := range opts {
		opt(cache)
	}

	if len(cache.name) == 0 {
		cache.name = defaultCacheName
	}
	cache.stats = newCacheStat(cache.name, cache.size)

	timingWheel, err := NewTimingWheel(time.Second, slots, func(k, v any) {
		key, ok := k.(string)
		if !ok {
			return
		}

		cache.Del(key)
	})
	if err != nil {
		return nil, err
	}

	cache.timingWheel = timingWheel
	return cache, nil
}

// Del deletes the item with the given key from c.
func (c *Cache[T]) Del(key string) {
	c.lock.Lock()
	delete(c.data, key)
	c.lruCache.remove(key)
	c.lock.Unlock()
	c.timingWheel.RemoveTimer(key)
}

// Get returns the item with the given key from c.
func (c *Cache[T]) Get(key string) (any, bool) {
	value, ok := c.doGet(key)
	if ok {
		c.stats.IncrementHit()
	} else {
		c.stats.IncrementMiss()
	}

	return value, ok
}

// Set sets value into c with key.
func (c *Cache[T]) Set(key string, value T) {
	c.SetWithExpire(key, value, c.expire)
}

// SetWithExpire sets value into c with key and expire with the given value.
func (c *Cache[T]) SetWithExpire(key string, value T, expire time.Duration) {
	c.lock.Lock()
	_, ok := c.data[key]
	c.data[key] = value
	c.lruCache.add(key)
	c.lock.Unlock()

	expiry := c.unstableExpiry.AroundDuration(expire)
	if ok {
		c.timingWheel.MoveTimer(key, expiry)
	} else {
		c.timingWheel.SetTimer(key, value, expiry)
	}
}

// Take returns the item with the given key.
// If the item is in c, return it directly.
// If not, use fetch method to get the item, set into c and return it.
func (c *Cache[T]) Take(key string, fetch func() (T, error)) (T, error) {
	if val, ok := c.doGet(key); ok {
		c.stats.IncrementHit()
		return val, nil
	}

	var fresh bool
	val, err := c.barrier.Do(key, func() (T, error) {
		// because O(1) on map search in memory, and fetch is an IO query,
		// so we do double-check, cache might be taken by another call
		if val, ok := c.doGet(key); ok {
			return val, nil
		}

		v, e := fetch()
		if e != nil {
			var def T
			return def, e
		}

		fresh = true
		c.Set(key, v)
		return v, nil
	})
	if err != nil {
		var def T
		return def, err
	}

	if fresh {
		c.stats.IncrementMiss()
		return val, nil
	}

	// got the result from previous ongoing query
	c.stats.IncrementHit()
	return val, nil
}

func (c *Cache[T]) doGet(key string) (T, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, ok := c.data[key]
	if ok {
		c.lruCache.add(key)
	}

	return value, ok
}

func (c *Cache[T]) onEvict(key string) {
	// already locked
	delete(c.data, key)
	c.timingWheel.RemoveTimer(key)
}

func (c *Cache[T]) size() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return len(c.data)
}

// WithLimit customizes a Cache with items up to limit.
func WithLimit[T any](limit int) CacheOption[T] {
	return func(cache *Cache[T]) {
		if limit > 0 {
			cache.lruCache = newKeyLru(limit, cache.onEvict)
		}
	}
}

// WithName customizes a Cache with the given name.
func WithName[T any](name string) CacheOption[T] {
	return func(cache *Cache[T]) {
		cache.name = name
	}
}

type (
	lru interface {
		add(key string)
		remove(key string)
	}

	emptyLru struct{}

	keyLru struct {
		limit    int
		evicts   *list.List
		elements map[string]*list.Element
		onEvict  func(key string)
	}
)

func (elru emptyLru) add(string) {
}

func (elru emptyLru) remove(string) {
}

func newKeyLru(limit int, onEvict func(key string)) *keyLru {
	return &keyLru{
		limit:    limit,
		evicts:   list.New(),
		elements: make(map[string]*list.Element),
		onEvict:  onEvict,
	}
}

func (klru *keyLru) add(key string) {
	if elem, ok := klru.elements[key]; ok {
		klru.evicts.MoveToFront(elem)
		return
	}

	// Add new item
	elem := klru.evicts.PushFront(key)
	klru.elements[key] = elem

	// Verify size not exceeded
	if klru.evicts.Len() > klru.limit {
		klru.removeOldest()
	}
}

func (klru *keyLru) remove(key string) {
	if elem, ok := klru.elements[key]; ok {
		klru.removeElement(elem)
	}
}

func (klru *keyLru) removeOldest() {
	elem := klru.evicts.Back()
	if elem != nil {
		klru.removeElement(elem)
	}
}

func (klru *keyLru) removeElement(e *list.Element) {
	klru.evicts.Remove(e)
	key := e.Value.(string)
	delete(klru.elements, key)
	klru.onEvict(key)
}

type cacheStat struct {
	name         string
	hit          uint64
	miss         uint64
	sizeCallback func() int
}

func newCacheStat(name string, sizeCallback func() int) *cacheStat {
	st := &cacheStat{
		name:         name,
		sizeCallback: sizeCallback,
	}
	go st.statLoop()
	return st
}

func (cs *cacheStat) IncrementHit() {
	atomic.AddUint64(&cs.hit, 1)
}

func (cs *cacheStat) IncrementMiss() {
	atomic.AddUint64(&cs.miss, 1)
}

func (cs *cacheStat) statLoop() {
	ticker := time.NewTicker(statInterval)
	defer ticker.Stop()

	for range ticker.C {
		hit := atomic.SwapUint64(&cs.hit, 0)
		miss := atomic.SwapUint64(&cs.miss, 0)
		total := hit + miss
		if total == 0 {
			continue
		}
		percent := 100 * float32(hit) / float32(total)
		xlog.Statf("cache(%s) - qpm: %d, hit_ratio: %.1f%%, elements: %d, hit: %d, miss: %d",
			cs.name, total, percent, cs.sizeCallback(), hit, miss)
	}
}
