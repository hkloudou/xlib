package xsort

import (
	"sort"
)

func Int64(slice []int64) {
	sort.Sort(Int64Slice(slice))
}

func Uint64(slice []int64) {
	sort.Sort(Uint64Slice(slice))
}

func Int32(slice []int32) {
	sort.Sort(Int32Slice(slice))
}

func Uint32(slice []uint32) {
	sort.Sort(Uint32Slice(slice))
}

//Uint64Slice Uint64Slice
type Uint64Slice []int64

func (p Uint64Slice) Len() int {
	return len(p)
}

func (p Uint64Slice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Uint64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//Int64Slice Int64Slice
type Int64Slice []int64

func (p Int64Slice) Len() int {
	return len(p)
}

func (p Int64Slice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Int64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//Int32Slice Int32Slice
type Int32Slice []int32

func (p Int32Slice) Len() int {
	return len(p)
}

func (p Int32Slice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Int32Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//Uint32Slice Uint32Slice
type Uint32Slice []uint32

func (p Uint32Slice) Len() int {
	return len(p)
}

func (p Uint32Slice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Uint32Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
