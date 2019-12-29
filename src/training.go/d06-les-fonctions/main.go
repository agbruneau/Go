package main

import (
	"fmt"
)

func printInfoNoParam() {
	fmt.Printf("Name=%s, age=%d, email=%s\n", "Bob", 10, "bob@golang.org")
}

func printInfoParams(name string, age int, email string) {
	fmt.Printf("Name=%s, age=%d, email=%s\n", name, age, email)
}

// If sequential parameters share the same type,
// the type can be specified in last position
func avg(x, y float64) float64 {
	return (x + y) / 2
}

// Variable sum is declared in the signature of the function
func sumNamedReturn(x, y, z int) (sum int) {
	sum = x + y + z
	return
}

func main() {
	printInfoNoParam()
	printInfoParams("Alice", 15, "alice@golang.org")

	result := avg(16.3, 25.0)
	fmt.Printf("Average result=%f\n", result)

	sum := sumNamedReturn(10, 25, 7)
	fmt.Printf("Sum result=%d\n", sum)
}
