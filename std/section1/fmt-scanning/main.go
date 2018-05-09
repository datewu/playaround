package main

import (
	"fmt"
	"strings"
)

func main() {
	scanWithSscan()
	scanWithSscanf()
	scanWithSscanln()

	//scanWithScan()
	scanFromReader()

	customFormatter()
	scanWithCustomScanner()

}

func scanWithSscan() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)
	if _, err := fmt.Sscan(" \r\n \t  5  7  9sdfsf", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Sscan: %d, %d, %s\n", d1, d2, s1)
}

func scanWithSscanf() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)
	if _, err := fmt.Sscanf(" \r \t  5  7 ,  9sdfsf", "%d %d , %s", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Sscanf: %d, %d, %s\n", d1, d2, s1)
}

func scanWithSscanln() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)
	if _, err := fmt.Sscanln(" \r \t  5  7   9sd\nfsf", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Sscanln: %d, %d, %s\n", d1, d2, s1)
}

func scanWithScan() {
	var s1, s2, s3 string
	fmt.Print("You inputs please: ")
	if _, err := fmt.Scan(&s1, &s2, &s3); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Scan: %s, %s, %s\n", s1, s2, s3)
}

func scanFromReader() {
	var s1, s2, s3 string
	r := strings.NewReader("A N C")
	fmt.Fscan(r, &s1, &s2, &s3)
	fmt.Printf("After scanning with Fscan: %s, %s, %s\n", s1, s2, s3)
}

type coords struct {
	x, y, z int
}

func (c *coords) Scan(f fmt.ScanState, v rune) error {
	switch v {
	case '$':
		n, err := fmt.Fscanf(f, "(%v, %v, %v)", &c.x, &c.y, &c.z)
		if err != nil {
			return err
		}
		if n != 3 {
			return fmt.Errorf("3 values expected, got %d", n)
		}
	}
	return nil
}

func (c coords) Format(f fmt.State, v rune) {
	switch v {
	case '$':
		fmt.Fprintf(f, "$(%v, %v, %v)", c.x, c.y, c.z)
	case 'v':
		fmt.Fprintf(f, "Coords{x=%v, y=%v, z=%v}", c.x, c.y, c.z)
	default:
		fmt.Fprintf(f, "%v, %v, %v", c.x, c.y, c.z)
	}
}

func scanWithCustomScanner() {
	var c coords
	if _, err := fmt.Sscanf("(2, 9, 10)", "%$", &c); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", c)
	fmt.Printf("%$\n", c)
}

func customFormatter() {
	c := coords{998, 990, 976}
	fmt.Printf("%$\n", c)
	fmt.Printf("%v\n", c)
	fmt.Printf("%d\n", c)
}
