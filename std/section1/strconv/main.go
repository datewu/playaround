package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	stringToInt()
	intToString()

	stringToAltBaseInt()
	formatInt()
	appendInt()

	parseFloat()
	formatFloat()
	parseBool()

	quoting()
}

func stringToInt() {
	n, err := strconv.Atoi("3454")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%v - %T\n", n, n)
}

func intToString() {
	n := 2335
	//n := 23434000000000000
	s := strconv.Itoa(n)
	fmt.Printf("%v - %T\n", s, s)
}

func stringToAltBaseInt() {
	s := "1604"
	n, err := strconv.ParseInt(s, 8, 0)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%v - %T\n", n, n)
	fmt.Printf("%o\n", n)
}

func formatInt() {
	x := int64(234)
	s := strconv.FormatInt(x, 10)
	fmt.Printf("%T - %v\n", s, s)

	x = int64(-234)
	s = strconv.FormatInt(x, 2)
	fmt.Printf("%T - %v\n", s, s)
}

func appendInt() {
	b := []byte("what are we appending?: ")
	b = strconv.AppendInt(b, -234, 10)
	fmt.Printf("%v\n", string(b))
}

func parseFloat() {
	s := "3.141592654"
	f, err := strconv.ParseFloat(s, 64)
	fmt.Printf("%v - %T - %v\n", f, f, err)
}

func formatFloat() {
	f := 3.141592654
	s := strconv.FormatFloat(f, 'e', 2, 64)
	fmt.Printf("%v - %T\n", s, s)
}

func parseBool() {
	payload := []string{
		"1", "t", "T", "TRUE", "true", "yes", "Yes", "Yay", "Yippie", "0", "f", "F", "FALSE", "False", "Nope", "Nah",
		"No",
	}
	for _, s := range payload {
		b, err := strconv.ParseBool(s)
		fmt.Printf("%v - %T - %v\n", b, b, err)
	}
}

func quoting() {
	fmt.Println(strconv.Quote("There is a ∂ symbol here as well as a tab 	."))
	fmt.Println(strconv.QuoteToASCII("There is a ∂ symbol here as well as a tab 	."))
}
