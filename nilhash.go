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
	h["dota"] = 999
	fmt.Println(h)
	h = nil
	//h["dota"] = 996
	fmt.Println(h)
	json.Unmarshal([]byte(j), &h)
	json.Unmarshal([]byte(k), &h)
	json.Unmarshal([]byte(m), &h)
	h["dota"] = 996
	fmt.Println(h)
	fmt.Println("++++++++++++++")

	fmt.Println(graph)
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
