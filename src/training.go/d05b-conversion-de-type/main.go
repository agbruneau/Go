package main

import (
	"fmt"
)

func main() {
	var percentage float64 = 2.0

	// 34 is upcasted to float, percentage is still a float64
	percentage += 34
	fmt.Printf("Current percentage %f%%\n", percentage)
	fmt.Printf("Int value %d%%\n", int(percentage))

	n := 42
	f := float64(n) + .42
	fmt.Printf("float=%f\n", f)
}
