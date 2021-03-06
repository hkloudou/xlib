package xface

import "time"

type Cache[T any] interface {
	Get(key string) (*T, error)
	Set(key string, ttl time.Duration, obj *T) error
	Del(key ...string) error
	Validator(func(obj *T) error) Cache[T]
}
