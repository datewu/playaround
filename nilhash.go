package main

import (
	"encoding/json"
	"fmt"
)

var j = `{ "lol": 8 }`
var k = `{ "dota": 10 }`
var m = `{ "cf": 12 }`

var graph = make(map[string]map[string]bool)

var d map[string]bool

func main() {
	//var h map[string]int
	h := make(map[string]int)

	h["pop"] = 3
	fmt.Println(h)
	//h = nil
	//h["dota"] = 999
	json.Unmarshal([]byte(j), &h)
	json.Unmarshal([]byte(k), &h)
	json.Unmarshal([]byte(m), &h)
	fmt.Println(h)
	fmt.Println(graph["lol"])
	fmt.Println(graph["lol"]["dd"])
	fmt.Println(d["ddd"])
}

/*
func main() {
	var h map[string]int
	f, _ := os.Open("./test.json")
	json.NewDecoder(f).Decode(&h)
	fmt.Println(h)
}
*/
