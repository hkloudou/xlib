package xface

import (
	"fmt"
	"time"
)

//cacherSecondary 二级缓存系统
type secondary[T any] struct {
	l1          Cache[T] //l1 cacher
	l2          Cache[T] //l2 cacher
	l1_Duration time.Duration
	l2_Duration time.Duration
}

//NewSecondaryCacher 新建一个二级缓存系统
func NewSecondaryCacher[T any](
	l1 Cache[T], l2 Cache[T],
	l1_Duration, l2_Duration time.Duration,
) Cache[T] {
	return &secondary[T]{
		l1:          l1,
		l2:          l2,
		l1_Duration: l1_Duration,
		l2_Duration: l2_Duration,
	}
}

// Get	1、先从L1读取	2、失败则从L2读取，并更新L1
func (m *secondary[T]) Get(key string) (*T, error) {
	if tmp, err := m.l1.Get(key); err == nil {
		return tmp, nil
	} else if tmp, err := m.l2.Get(key); err == nil {
		m.l1.Set(key, m.l1_Duration, tmp) //这里不判断返回错误，可以容错L1 故障（L1一般内存，虽然概率较低）
		return tmp, nil
	}
	return nil, fmt.Errorf("record not found")
}

//Set 1、先设置L2	2、然后删除L1
func (m *secondary[T]) Set(key string, ttl time.Duration, obj *T) error {
	if ttl == -1 {
		ttl = m.l2_Duration
	}
	//强制要求L2的可用性
	if err := m.l2.Set(key, ttl, obj); err != nil {
		return err
	} else if err := m.l1.Del(key); err != nil {
		return err
	}
	return nil
}

// DEL	1、先删除L1	2、再删除L2	3、最后删除L1
func (m *secondary[T]) Del(key ...string) error {
	// if err := m.l1.Del(key...); err != nil {
	// 	return err
	// } else
	//考虑到L1存在内存，L2一般是REDIS，先优先删除L2把，避免在删除L2后又被GET重新XIERU
	if err := m.l2.Del(key...); err != nil {
		return err
	} else if err := m.l1.Del(key...); err != nil {
		return err
	}
	return nil
}

func (m *secondary[T]) Validator(fc func(obj *T) error) Cache[T] {
	m.l1.Validator(fc)
	m.l2.Validator(fc)
	return m
}
