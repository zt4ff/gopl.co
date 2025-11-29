package ch4

import (
	"crypto/sha256"
	"fmt"
)

func SHA256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("y"))

	a := []rune("welcome")
	fmt.Println(len(a))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

// this function zeroes the content of a [32]byte array
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
