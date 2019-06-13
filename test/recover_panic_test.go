package main

import (
	"testing"
	"fmt"
)

func recoverPanic(fn func(*testing.T)) func(*testing.T) {
	return func(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("there is a panic in testing:", r)
		}
	}()
	fn(t)
	}

}

func TestResourceManager(m *testing.T) {
	m.Run("env", recoverPanic(testPanic))
}

func testPanic(t *testing.T) {
	panic("lol")
}
