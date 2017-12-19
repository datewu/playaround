package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	verbose = flag.Bool("v", false, "show verbosr progress messages")
	sema    = make(chan struct{}, 2000)
	done    = make(chan struct{})
)

func walkDir(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() {
		<-sema
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "du: ", err)
		return nil
	}
	return entries
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	fileSizes := make(chan int64)

	var wg sync.WaitGroup
	wg.Add(len(roots))
	go func() {
		for _, root := range roots {
			go walkDir(root, &wg, fileSizes)
		}

		go func() {
			wg.Wait()
			close(fileSizes)
		}()
	}()

	display(fileSizes)

}

func display(fileSize <-chan int64) {
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(499 * time.Microsecond)
	}

	var nfiles, nbytes int64

loop:
	for {
		select {
		case <-done:
			for range fileSize {
			}
			return
		case size, ok := <-fileSize:
			if !ok {
				break loop // fileSize was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}
