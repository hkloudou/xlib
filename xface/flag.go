package xface

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/hkloudou/xlib/xflag"
)

var _flags = sync.Map{}
var _list = []string{}
var _lk = &sync.RWMutex{} //主要对_list

type FlagConfig[T any] interface {
	Flags() []xflag.Flag
	Action(c *xflag.Context) error
	Instance() *T
}

func FlagRegister[T any](name string, obj FlagConfig[T]) {
	_lk.Lock()
	defer _lk.Unlock()
	_flags.Store(name, obj)
	if !strings.Contains(","+strings.Join(_list, ",")+",", name) {
		_list = append(_list, name)
	}
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

//FlagAction 让Action 有序执行，这样在一些场景下实现智能依赖
func FlagAction(c *xflag.Context) (err error) {
	_lk.RLock()
	defer _lk.RUnlock()
	for i := 0; i < len(_list); i++ {
		name := _list[i]
		if value, found := _flags.Load(name); !found {
			continue
		} else {
			mt := reflect.ValueOf(value).MethodByName("Action")
			res := mt.Call([]reflect.Value{reflect.ValueOf(c)})
			if !res[0].IsNil() {
				err = res[0].Interface().(error)
				if err != nil {
					return
				}
			}
		}
	}
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
