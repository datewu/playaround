package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkUnbufferedWriter(b *testing.B) {
	performWrite(b, temFileOrFatal())
}

func BenchmarkBufferedWriter(b *testing.B) {
	bufferedFile := bufio.NewWriter(temFileOrFatal())
	performWrite(b, bufferedFile)
}

func temFileOrFatal() *os.File {
	file, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return file
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()

	for bt := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{bt.(byte)})
	}

}
