package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

// Pool manages a set of resources that can be shared safely by
// multiple goroutines. The resource being managed must implement
// the io.Closer interface.
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	hold      chan struct{}
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed is returned when an Acquire returns on a closed pool.
var ErrPoolClosed = errors.New("Pool has been closed")

// New creates a pool that manages resources. A pool requires a function
// that can allocate a new resource and the size of the pool.
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
		hold:      make(chan struct{}, size),
	}, nil
}

// Acquire retrieves a resource from the pool.
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		fmt.Println("Acquire:", "Share Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	case p.hold <- struct{}{}:
		fmt.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release places a new resource onto the pool.
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		fmt.Println("Release:", "to Queue")
	default:
		fmt.Println("Release:", "Closing, the Queue seems to be full")
		r.Close()
	}
}

// Close will shutdown the pool and close all existing resources.
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	p.closed = true

	close(p.resources)
}
