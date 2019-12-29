package main

import (
	"fmt"
)

func start() {
	fmt.Println("Start")
}

func finish() {
	fmt.Println("Finish")
}

func main() {
	// defer is added in a stack
	// and executed in a LIFO order
	start()
	defer finish()

	names := []string{"Bob", "Alice", "Bobette", "John"}
	for _, n := range names {
		fmt.Printf("Hello %v...\n", n)
		defer fmt.Printf("Goodbye %v\n", n)
	}
}
