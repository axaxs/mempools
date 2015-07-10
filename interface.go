package mempools


type InterfacePool struct {
	list []interface{}
	index int
	New func() interface{}
}

func (o *InterfacePool) Alloc() interface{} {
        if o.index >= len(o.list) {
                o.list = append(o.list, o.New())
        }
        o.index++
        return o.list[o.index-1]
}

func (o *InterfacePool) Reset() {
	o.index = 0
}

func (o *InterfacePool) Destroy() {
	o = &InterfacePool{}
}
