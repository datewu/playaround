package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("need two paramters, but got", len(os.Args)-1)
	}
	in := os.Args[1]
	salt := os.Args[2]
	add(in, in+".new", salt+"x")

	fmt.Println("good luck")
}

func add(in, out, salt string) (err error) {
	source, err := os.Open(in)
	if err != nil {
		log.Println(err)
		return err
	}
	defer source.Close()

	result, err := os.Create(out)
	if err != nil {
		log.Println(err)
		return err
	}
	defer result.Close()

	w := bufio.NewWriter(result)
	scanner := bufio.NewScanner(source)

	scanner.Split(bufio.ScanLines)
	scanner.Scan() // skip the first line
	_, err = w.WriteString(scanner.Text() + "\n")
	for scanner.Scan() {
		lineString := scanner.Text()
		if lineString[0] != '<' {
			lineString = salt + lineString
		}
		_, err = w.WriteString(lineString + "\n")
		if err != nil {
			return
		}
	}
	err = scanner.Err()
	return
}
