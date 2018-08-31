package main

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type a struct {
	A string `json:"direct"`
	B *bob   `json:" ,inline"`
}

type bob struct {
	Name string
	Age  int
}

func main() {
	test := a{"one", &bob{"dota", 99}}

	json.NewEncoder(os.Stdout).Encode(&test)

	s := `
	{"direct":"one"," ":{"Name":"dota","Age":99}}
	`
	reader := bytes.NewBufferString(s)

	var xxx a
	json.NewDecoder(reader).Decode(&xxx)

	spew.Dump(xxx)

}
