package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	idxBits           = 6
	idxMask           = 1<<idxBits - 1
	asscalCollections = "abcdefghijklmnopqrstuvwxyzsABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!"
	idxMax            = 63 / idxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func main() {
	a := make([]string, 0)
	a = append(a, randSting(6))
	a = append(a, randSting(4))
	a = append(a, randSting(4))
	a = append(a, randSting(6))
	a = append(a, randSting(8))
	a = append(a, randSting(8))
	fmt.Printf("%q\n", a)

}

func randSting(length int) string {
	b := make([]byte, length)

	for i, cache, remain := length-1, src.Int63(), idxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), idxMax
		}

		if idx := int(idxMask & cache); idx < len(asscalCollections) {
			b[i] = asscalCollections[idx]
			i--
		}
		cache >>= idxBits
		remain--
	}
	return string(b)
}
