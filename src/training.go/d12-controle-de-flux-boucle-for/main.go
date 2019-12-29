package main

import "fmt"

func main() {

	// for iteration loop
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Sum result=%d\n", sum)

	// while loop
	eventCnt := 0
	for eventCnt < 3 {
		fmt.Println("Retrieving events...")
		eventCnt++
		if eventCnt == 3 {
			fmt.Printf("Got %d notifications, updating Phone notifications\n", eventCnt)
		}
	}

	// forever loop
	i := 0
	for {
		i++
		if i%2 != 0 {
			fmt.Println("Odd looping")
			continue
		}
		fmt.Println("Looping...")

		if i >= 10 {
			fmt.Println("Loop end")
			break
		}
	}

}
