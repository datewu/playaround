package main

import (
	"fmt"
)

func main() {
	var m map[string]bool = nil
//	var m map[string]bool
//	m := map[string]bool{"lll": true}
//	m["ddd"] = true
	delete(m, "foo")
	fmt.Println("DONE")

}
