package xface

import (
	"fmt"
	"time"
)

type cacherMixer[T any] struct {
	c1 Cacher[T] //l1 cacher
	c2 Cacher[T] //l2 cacher
}

func NewCacheMixer[T any](c1 Cacher[T], c2 Cacher[T]) Cacher[T] {
	return &cacherMixer[T]{
		c1: c1,
		c2: c2,
	}
}

func (m *cacherMixer[T]) Get(key string) (*T, error) {
	if tmp, err := m.c1.Get(key); err == nil {
		return tmp, nil
	} else if tmp, err := m.c2.Get(key); err == nil {
		m.c1.Set(key, 0, tmp)
		return tmp, nil
	}
	return nil, fmt.Errorf("not found")
}

func (m *cacherMixer[T]) Set(key string, ttl time.Duration, obj *T) error {
	m.c2.Set(key, ttl, obj)
	m.c1.Set(key, ttl, obj)
	return nil
}

func (m *cacherMixer[T]) Del(key ...string) error {
	m.c2.Del(key...)
	m.c1.Del(key...)
	return nil
}

func (m *cacherMixer[T]) Validator(fc func(obj *T) error) {
	m.c1.Validator(fc)
	m.c2.Validator(fc)
}
