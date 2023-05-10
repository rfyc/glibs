package std

import (
	"sync"
)

type Chan struct {
	ch      chan T
	rwmutex *sync.RWMutex
	closed  bool
}

func NewChan(size ...int) *Chan {
	return &Chan{
		ch:      make(chan T, append(size, 0)[0]),
		rwmutex: &sync.RWMutex{},
	}
}

func (this *Chan) Recv() chan T {
	return this.ch
}

func (this *Chan) Send(data any) bool {
	this.rwmutex.RLock()
	defer this.rwmutex.RUnlock()
	if this.closed {
		return false
	} else {
		this.ch <- NewT(data)
		return true
	}
}

func (this *Chan) Close() {
	this.rwmutex.Lock()
	defer this.rwmutex.Unlock()
	close(this.ch)
	this.closed = true
}

func (this *Chan) IsClosed() bool {
	return this.closed
}
