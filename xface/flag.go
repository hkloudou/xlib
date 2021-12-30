package xface

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/hkloudou/xlib/xflag"
)

var _flags = sync.Map{}

type FlagConfig[T any] interface {
	Flags() []xflag.Flag
	Action(c *xflag.Context) error
	Instance() *T
}

func FlagRegister[T any](name string, obj FlagConfig[T]) {
	_flags.Store(name, obj)
}

func FlagGet[T any](name string) (*T, error) {
	if item, found := _flags.Load(name); !found {
		return nil, fmt.Errorf("not found")
	} else if obj, ok := item.(FlagConfig[T]); !ok {
		return nil, fmt.Errorf("not type equal")
	} else {
		return obj.Instance(), nil
	}
}

func FlagAction(c *xflag.Context) (err error) {
	_flags.Range(func(key, value any) bool {
		mt := reflect.ValueOf(value).MethodByName("Action")
		res := mt.Call([]reflect.Value{reflect.ValueOf(c)})
		if !res[0].IsNil() {
			err = res[0].Interface().(error)
			return false
		}
		return true
	})
	return
}

func Flags() []xflag.Flag {
	var tmp = make([]xflag.Flag, 0)
	_flags.Range(func(key, value any) bool {
		mt := reflect.ValueOf(value).MethodByName("Flags")
		res := mt.Call(nil)
		tmp = append(tmp, res[0].Interface().([]xflag.Flag)...)
		return true
	})
	return tmp
}
