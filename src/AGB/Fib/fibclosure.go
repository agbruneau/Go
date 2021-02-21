package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second := 1, 0
	return func() int {
		first, second = second, first+second
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100000000; i++ {
		fmt.Println(f())
	}
}