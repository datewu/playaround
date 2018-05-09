package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// name := flag.String("name", "Gopher", "name of the gooher")
	// age := flag.Int("age", 2, "age of the gooher")
	// shy := flag.Bool("shy", false, "is the gopher shy?")
	// flag.Parse()
	// fmt.Printf("Gopher Stats\nName: %s\nAge: %d\nShy: %t\n", *name, *age, *shy)

	var (
		name string
		age  int
		shy  bool
	)
	flag.StringVar(&name, "name", defaultName(), "name of the gooher")
	flag.IntVar(&age, "age", defaultAge(), "age of the gooher")
	flag.BoolVar(&shy, "shy", defaultShyness(), "is the gopher shy?")
	flag.Parse()
	fmt.Printf("Gopher Stats\nName: %s\nAge: %d\nShy: %t\n", name, age, shy)
}

func defaultName() string {
	if d := os.Getenv("GOPHER_DEFAULT_NAME"); d != "" {
		return d
	}
	return "Gopher"
}

func defaultAge() int {
	if d := os.Getenv("GOPHER_DEFAULT_AGE"); d != "" {
		if age, err := strconv.Atoi(d); err == nil {
			return age
		}
	}
	return 2
}

func defaultShyness() bool {
	if d := os.Getenv("GOPHER_DEFAULT_SHYNESS"); d != "" {
		if shy, err := strconv.ParseBool(d); err == nil {
			return shy
		}
	}
	return true
}
