package main

import (
	"fmt"
)

func main() {
	var names [3]string
	fmt.Printf("names=%v (len=%d)\n", names, len(names))

	names[0] = "Bob"
	names[2] = "Alice"
	fmt.Printf("names=%v\n", names)
	fmt.Printf("name[2]=%v\n", names[2])

	// declare with all values and shortcut :=
	odds := [4]int{1, 3, 5, 7}
	fmt.Printf("odds=%v (len=%d)\n", odds, len(odds))
}
