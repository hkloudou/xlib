package xarray

import "testing"

func Test_List(t *testing.T) {
	l := NewList[int](false)

	t.Log(l.Push(11))
	t.Log(l.Push(12))
	t.Log(l.Push(13))
	t.Log(l.l)
	l.Shift()
	t.Log(l.l)
	l.UnShift(1, 2, 12)
	t.Log(l.l)
	t.Log(l.Distant())
	// t.Log(l.Pop())
	// t.Log(l.l)
	// t.Log(l.Pop())
	// t.Log(l.l)
	// t.Log(l.Pop())
	// t.Log(l.l)
}
