package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"./pool"
)

const (
	maxGoroutines   = 50
	pooledResources = 20
)

type dbConnection struct{ ID int32 }

// Close implements the io.Closer intrface so dbConnection
// can be managed by the pool. Close performs any resource
// release management.
func (d *dbConnection) Close() error {
	log.Println("Close: Connection:", d.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Fatalln(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()

	fmt.Println("Shutdown Pool.")
	p.Close()
	fmt.Println("Done")
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
