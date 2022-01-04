package xface

import (
	"fmt"
	"time"
)

//cacherSecondary 二级缓存系统
type cacherSecondary[T any] struct {
	l1          Cacher[T] //l1 cacher
	l2          Cacher[T] //l2 cacher
	l1_Duration time.Duration
	l2_Duration time.Duration
}

//NewSecondaryCacher 新建一个二级缓存系统
func NewSecondaryCacher[T any](
	l1 Cacher[T], l2 Cacher[T],
	l1_Duration, l2_Duration time.Duration,
) Cacher[T] {
	return &cacherSecondary[T]{
		l1:          l1,
		l2:          l2,
		l1_Duration: l1_Duration,
		l2_Duration: l2_Duration,
	}
}

// Get
// 1、先从L1读取
// 2、失败则从L2读取，并更新L1
func (m *cacherSecondary[T]) Get(key string) (*T, error) {
	if tmp, err := m.l1.Get(key); err == nil {
		return tmp, nil
	} else if tmp, err := m.l2.Get(key); err == nil {
		m.l1.Set(key, m.l1_Duration, tmp)
		return tmp, nil
	}
	return nil, fmt.Errorf("record not found")
}

// Set
// 1、先设置L2
// 2、然后删除L1
func (m *cacherSecondary[T]) Set(key string, ttl time.Duration, obj *T) error {
	if err := m.l2.Set(key, ttl, obj); err != nil {
		return err
	} else if err := m.l1.Del(key); err != nil {
		return err
	}
	return nil
}

func (m *cacherSecondary[T]) Del(key ...string) error {
	if err := m.l1.Del(key...); err != nil {
		return err
	} else if m.l2.Del(key...); err != nil {
		return err
	}
	return nil
}

func (m *cacherSecondary[T]) Validator(fc func(obj *T) error) Cacher[T] {
	m.l1.Validator(fc)
	m.l2.Validator(fc)
	return m
}
