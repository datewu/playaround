/*
create role lol with password 'dota' login;

create database yace;

create table data (
	id serial primary key,
	raw bytea);


// check table size
SELECT nspname || '.' || relname AS "relation",
    pg_size_pretty(pg_relation_size(C.oid)) AS "size"
  FROM pg_class C
  LEFT JOIN pg_namespace N ON (N.oid = C.relnamespace)
  WHERE nspname NOT IN ('pg_catalog', 'information_schema')
  ORDER BY pg_relation_size(C.oid) DESC
  LIMIT 20;
*/
package main

import (
	"database/sql"
	"flag"
	"log"
	"sync"
	"sync/atomic"
	"time"

	_ "github.com/lib/pq"
)

var (
	runfor     = flag.Duration("r", 30*time.Second, "duration of time to run")
	con        = flag.Int("c", 50, "concurrent")
	tcpAddress = flag.String("address", "127.0.0.1:5432", "<addr>:<port> to connect to nsqd")
	size       = flag.Int("size", 10, "size of messages")
)

var totalMsgCount int64

func writePG(td time.Duration, db *sql.Stmt, data []byte, rdyChan chan int, goChan chan int) {
	rdyChan <- 1
	<-goChan
	var msgCount int64
	endTime := time.Now().Add(td)
	for {
		msgCount++
		db.Exec(string(data))
		if time.Now().After(endTime) {
			break
		}
	}
	atomic.AddInt64(&totalMsgCount, msgCount)
}

func main() {
	flag.Parse()
	var wg sync.WaitGroup

	log.SetPrefix("[pg_writer] ")

	goChan := make(chan int)
	rdyChan := make(chan int)
	data := make([]byte, *size*1000)
	stmt, err := initInsert()
	if err != nil {
		log.Fatalln(err)
	}

	wg.Add(*con)
	for j := 0; j < *con; j++ {
		go func() {
			writePG(*runfor, stmt, data, rdyChan, goChan)
			wg.Done()
		}()
		<-rdyChan
	}
	start := time.Now()
	close(goChan)
	wg.Wait()
	end := time.Now()
	duration := end.Sub(start)
	tmc := atomic.LoadInt64(&totalMsgCount)
	log.Printf("duration: %s -  - %.03fops/s - %.03fus/op",
		duration,
		float64(tmc)/duration.Seconds(),
		float64(duration/time.Microsecond)/float64(tmc))
}

func initInsert() (stmt *sql.Stmt, err error) {
	uri := "postgres://lol:dota@" + *tcpAddress + "/yace?sslmode=disable"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return
	}
	stmt, err = db.Prepare("insert into data (raw) values ($1)")
	return
}
