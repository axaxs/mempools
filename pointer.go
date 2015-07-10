package mempools

import (
	"unsafe"
)

type Pointer unsafe.Pointer

type PointerPool struct {
	list  []Pointer
	index int
	New   func() Pointer
}

func (o *PointerPool) Alloc() Pointer {
	if o.index >= len(o.list) {
		o.list = append(o.list, Pointer(o.New()))
	}
	o.index++
	return o.list[o.index-1]
}

func (o *PointerPool) Reset() {
	o.index = 0
}

func (o *PointerPool) Destroy() {
	*o = PointerPool{New: o.New}
}
