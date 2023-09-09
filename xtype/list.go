package xtype

import (
	"encoding/json"
	"sync"
)

type List[T comparable] struct {
	mu *sync.RWMutex
	l  []T
}

func (l *List[T]) lock() {
	if l.mu != nil {
		l.mu.Lock()
	}
}

func (l *List[T]) unlock() {
	if l.mu != nil {
		l.mu.Unlock()
	}
}

func (l *List[T]) rLock() {
	if l.mu != nil {
		l.mu.RLock()
	}
}

func (l *List[T]) rUnlock() {
	if l.mu != nil {
		l.mu.RUnlock()
	}
}

// Concat The concat() method is used to merge two or more arrays. This method does not change the existing arrays, but instead returns a new array.
func (l *List[T]) Concat(items ...List[T]) *List[T] {
	// l.Push()
	for i := 0; i < len(items); i++ {
		items[i].ForEach(func(t T, i int) {
			l.Push(t)
		})
	}
	return l
}

// Entries The entries() method returns a new Array Iterator object that contains the key/value pairs for each index in the array.
func (l *List[T]) Entries() map[int]T {
	l.rLock()
	defer l.rUnlock()
	items := make(map[int]T, 0)
	for i := 0; i < len(l.l); i++ {
		items[i] = l.l[i]
	}
	return items
}

// Every The every() method tests whether all elements in the array pass the test implemented by the provided function. It returns a Boolean value.
func (l *List[T]) Every(fn func(T) bool) bool {
	l.rLock()
	defer l.rUnlock()
	for i := 0; i < len(l.l); i++ {
		if !fn(l.l[i]) {
			return false
		}
	}
	return true
}

// Fill The fill() method changes all elements in an array to a static value, from a start index (default 0) to an end index (default array.length). It returns the modified array.
func (l *List[T]) Fill(val T, poss ...int) *List[T] {
	// arr.fill(value[, start[, end]])
	l.lock()
	defer l.unlock()
	start := 0
	end := len(l.l)
	if len(poss) == 2 {
		start = poss[0]
		end = poss[1]
	} else if len(poss) == 1 {
		start = poss[0]
	}
	for i := 0; i < len(l.l); i++ {
		if i >= start && i <= end {
			l.l[i] = val
		}
	}
	return l
}

// Filter The filter() method creates a new array with all elements that pass the test implemented by the provided function.
func (l *List[T]) Filter(fn func(T) bool) *List[T] {
	l.rLock()
	defer l.rUnlock()
	lnew := NewList[T](l.mu != nil)
	for i := 0; i < len(l.l); i++ {
		if fn(l.l[i]) {
			lnew.l = append(lnew.l, l.l[i])
		}
	}
	return lnew
}

// Find The find() method returns the first element in the provided array that satisfies the provided testing function. If no values satisfy the testing function, undefined is returned.
func (l *List[T]) Find(fn func(T) bool) (T, bool) {
	l.rLock()
	defer l.rUnlock()
	for i := 0; i < len(l.l); i++ {
		if fn(l.l[i]) {
			return l.l[i], true
		}
	}
	var def T
	return def, false
}

// FindIndex The findIndex() method returns the index of the first element in the array that satisfies the provided testing function. Otherwise, it returns -1, indicating that no element passed the test.
func (l *List[T]) FindIndex(fn func(T) bool) int {
	l.rLock()
	defer l.rUnlock()
	for i := 0; i < len(l.l); i++ {
		if fn(l.l[i]) {
			return i
		}
	}
	return -1
}

// ForEach The forEach() method executes a provided function once for each array element.
func (l *List[T]) ForEach(fn func(T, int)) {
	l.rLock()
	defer l.rUnlock()
	for i := 0; i < len(l.l); i++ {
		fn(l.l[i], i)
	}
}

// Includes The includes() method determines whether an array includes a certain value among its entries, returning true or false as appropriate.
func (l *List[T]) Includes(items ...T) bool {
	l.rLock()
	defer l.rUnlock()
	if len(l.l) == 0 || len(items) == 0 {
		return false
	}

	for j := 0; j < len(items); j++ {
		found := false
		for i := 0; i < len(l.l); i++ {
			if items[j] == l.l[i] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// IndexOf The indexOf() method returns the first index at which a given element can be found in the array, or -1 if it is not present.
func (l *List[T]) IndexOf(val T) int {
	l.rLock()
	defer l.rUnlock()
	for i := 0; i < len(l.l); i++ {
		if l.l[i] == val {
			return i
		}
	}
	return -1
}

// LastIndexOf The lastIndexOf() method returns the last index at which a given element can be found in the array, or -1 if it is not present. The array is searched backwards, starting at fromIndex.
func (l *List[T]) LastIndexOf(val T) int {
	l.rLock()
	defer l.rUnlock()
	for i := len(l.l) - 1; i >= 0; i-- {
		if l.l[i] == val {
			return i
		}
	}
	return -1
}

// Map The map() method creates a new array populated with the results of calling a provided function on every element in the calling array.
func (l *List[T]) Map(fn func(val T) T) *List[T] {
	l.rLock()
	defer l.rUnlock()
	lnew := NewList[T](l.mu != nil)
	for i := 0; i < len(l.l); i++ {
		lnew.l = append(lnew.l, fn(l.l[i]))
	}
	return lnew
}

// Pop The pop() method removes the last element from an array and returns that element. This method changes the length of the array.
func (l *List[T]) Pop() (T, bool) {
	l.lock()
	defer l.unlock()
	if len(l.l) == 0 {
		var def T
		return def, false
	}
	item := l.l[len(l.l)-1]
	l.l = l.l[:len(l.l)-1]
	return item, true
}

// Push The push() method adds one or more elements to the end of an array and returns the new length of the array.
func (l *List[T]) Push(v ...T) int {
	l.lock()
	defer l.unlock()
	l.l = append(l.l, v...)
	return len(l.l)
}

// UnShift The unshift() method adds one or more elements to the beginning of an array and returns the new length of the array.
func (l *List[T]) UnShift(v ...T) int {
	l.lock()
	defer l.unlock()
	l.l = append(v, l.l...)
	return len(l.l)
}

// Shift The shift() method removes the first element from an array and returns that removed element. This method changes the length of the array.
func (l *List[T]) Shift() *List[T] {
	l.lock()
	defer l.unlock()
	if len(l.l) > 0 {
		l.remove(0)
	}
	return l
}

func (l *List[T]) insertAt(pos int, v T) {
	l.l, l.l[0] = append(l.l[:pos+1], l.l[pos:]...), v
	// return
}

func (l *List[T]) remove(i int) {
	l.l = l.l[:i+copy(l.l[i:], l.l[i+1:])]
	// return
}

func (l *List[T]) Some(fn func(val T) bool) bool {
	l.rLock()
	defer l.rUnlock()
	for i := 0; i < len(l.l); i++ {
		if fn(l.l[i]) {
			return true
		}
	}
	return false
}

func (l *List[T]) ToString() string {
	l.rLock()
	defer l.rUnlock()
	b, _ := json.Marshal(l.l)
	return string(b)
}

func (l *List[T]) ValueOf() []T {
	l.rLock()
	defer l.rUnlock()
	return l.l
}

// func (l *list[T]) Equals(rhs *list[T]) bool {
// 	if l.mu != nil {
// 		l.mu.RLock()
// 		defer l.mu.RUnlock()
// 	}
// 	if len(l.l) != len(rhs.l) {
// 		return false
// 	}
// 	for i := 0; i < len(l.l); i++ {
// 		if l.l[i] != rhs.l[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func (l *List[T]) Distant() *List[T] {
	l.rLock()
	defer l.rUnlock()
	ll := &List[T]{l: make([]T, 0)}
	if l.mu != nil {
		ll.mu = &sync.RWMutex{}
	}
	m := make(map[T]bool)
	for _, v := range l.l {
		if _, ok := m[v]; !ok {
			m[v] = true
			ll.l = append(ll.l, v)
		}
	}
	return ll
}

func (l *List[T]) Clone() *List[T] {
	l.rLock()
	defer l.rUnlock()
	ll := &List[T]{l: make([]T, len(l.l))}
	if l.mu != nil {
		ll.mu = &sync.RWMutex{}
	}
	copy(ll.l, l.l)
	return ll
}

func NewList[T comparable](safe bool) *List[T] {
	l := &List[T]{l: make([]T, 0)}
	if safe {
		l.mu = &sync.RWMutex{}
	}
	return l
}
