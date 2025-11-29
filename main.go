package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}

	b := a[:]

	fmt.Printf("Size: %v, Cap: %v\n", len(b), cap(b))

	b = append(b, 3, 3)

	fmt.Printf("Size: %v, Cap: %v\n", len(b), cap(b))

	fmt.Println(b)
}
