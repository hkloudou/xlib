package xface

import (
	"fmt"
	"sync"
	"time"
)

func NewMemoryCacher[T any]() Cacher[T] {
	return &memory[T]{_map: map[string]*T{}, _lock: sync.RWMutex{}, _validator: func(obj *T) error { return nil }}
}

type memory[T any] struct {
	_map       map[string]*T
	_lock      sync.RWMutex
	_validator func(obj *T) error
}

func (m *memory[T]) Get(key string) (*T, error) {
	m._lock.RLock()
	defer m._lock.RUnlock()
	if obj, found := m._map[key]; !found {
		return nil, fmt.Errorf("not found")
	} else if err := m._validator(obj); err != nil {
		return nil, err
	} else {
		return obj, nil
	}
}
func (m *memory[T]) Set(key string, ttl time.Duration, obj *T) error {
	if err := m._validator(obj); err != nil {
		return err
	}
	m._lock.Lock()
	defer m._lock.Unlock()
	m._map[key] = obj
	return nil
}
func (m *memory[T]) Del(keys ...string) error {
	m._lock.Lock()
	defer m._lock.Unlock()
	for i := 0; i < len(keys); i++ {
		delete(m._map, keys[i])
	}
	return nil
}

func (m *memory[T]) Validator(fc func(obj *T) error) {
	m._validator = fc
}
