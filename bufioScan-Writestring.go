package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("need two paramters, got", len(os.Args)-1)
	}
	in, salt := os.Args[1], os.Args[2]

	log.Println(add(in, in+".new", salt+"x"))
	fmt.Println("Good Luck")
}

func add(in, out, salt string) (err error) {
	source, err := os.Open(in)
	if err != nil {
		return
	}
	defer source.Close()

	result, err := os.Create(out)
	if err != nil {
		return
	}
	defer result.Close()

	w := bufio.NewWriter(result)
	scanner := bufio.NewScanner(source)

	scanner.Split(bufio.ScanLines)
	scanner.Scan() // skip the first line
	_, err = w.WriteString(scanner.Text() + "\n")

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] != '<' {
			line += salt
		}
		_, err = w.WriteString(line + "\n")
		if err != nil {
			return
		}
	}
	w.Flush()
	err = scanner.Err()
	return
}
