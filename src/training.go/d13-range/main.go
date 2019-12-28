package main

import (
	"fmt"
)

func main() {

	names := []string{"Bob", "Alice", "Bobette", "John"}
	for i, n := range names {
		fmt.Printf("Username=%s (index=%d)\n", n, i)
	}

	// range on string
	// Omit index / value
	for _, c := range "golang" {
		fmt.Printf("%v\n", string(c))
	}
}
