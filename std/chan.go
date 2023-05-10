package std

import (
	"sync"
)

type Chan struct {
	ch      chan any
	rwmutex *sync.RWMutex
	closed  bool
}

func NewChan(size ...int) *Chan {
	return &Chan{
		ch:      make(chan any, append(size, 0)[0]),
		rwmutex: &sync.RWMutex{},
	}
}

func (this *Chan) Recv() T {
	if v, ok := <-this.ch; ok {
		return NewT(v)
	}
	return T{}
}

func (this *Chan) Send(data any) bool {
	this.rwmutex.RLock()
	defer this.rwmutex.RUnlock()
	if this.closed {
		return false
	} else {
		this.ch <- data
		return true
	}
}

func (this *Chan) Close() {
	this.rwmutex.Lock()
	defer this.rwmutex.Unlock()
	close(this.ch)
	this.closed = true
}
