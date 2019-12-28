package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = -2
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	// When appending, the cap() of the underlying array can be * 2
	nums = append(nums, 64)
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	nums = append(nums, -8)
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	fmt.Println("----------------------")

	// Initialization with shortcut syntax
	letters := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))

	// subslicing
	sub1 := letters[:2]
	sub2 := letters[2:]
	fmt.Printf("%v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Println("----------------------")

	// Modifying underlying array
	sub2[0] = "UP"
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Println("----------------------")

	// Copy a slice to avoid modifying the same data
	// Keep in mind when reading a whole file and returning a slice...
	// ... The whole file is still in memory!
	subCopy := make([]string, len(sub1))
	copy(subCopy, sub1)
	subCopy[0] = "DOWN"
	fmt.Printf("COPY ==> %v (len=%d, cap=%d)\n", subCopy, len(subCopy), cap(subCopy))
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub2, len(sub1), cap(sub1))
}
