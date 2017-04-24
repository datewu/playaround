// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

import (
	"io/ioutil"
	"net/http"
	"sync"
)

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// NOTE: not concurrency-safe!
func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	e := m.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mu.Unlock()

		e.res.value, e.res.err = m.f(key)

		close(e.ready) //broadcase ready condition
	} else {
		// This is a repeat request for this key.
		m.mu.Unlock()
		<-e.ready // Wait for ready conditon
	}
	return e.res.value, e.res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Boby.Close()
	return ioutil.ReadAll(resp.Body)
}
