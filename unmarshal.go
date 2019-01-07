package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	s := `{"lol":"han bin ss"}`

	var dump struct{ Lol string }

	log.Println(json.Unmarshal([]byte(s), &dump))
	fmt.Printf("%#v", dump)
}
