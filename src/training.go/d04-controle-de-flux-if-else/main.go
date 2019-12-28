package main

import (
	"fmt"
)

func main() {

	day := 05 // day of the month [01, 31]

	if day < 15 {
		fmt.Printf("We're in the first half of the month (day=%d)\n", day)
	} else if day == 18 {
		fmt.Printf("We're on the special day (day=%d)\n", day)
	} else {
		fmt.Printf("We're in the second half of the month (day=%d)\n", day)
	}

	year, month, day := 2009, 11, 10
	fmt.Printf("Date=%d/%d/%d\n", year, month, day)

	if year == 2009 && month == 11 && day == 10 {
		fmt.Println("This is the first release of Go !")
	} else if year == 2009 || month == 11 || day == 10 {
		fmt.Println("At least on part is right... :/")
	} else {
		fmt.Println("Just another day ")
	}

	if count := 12; count > 10 {
		fmt.Printf("We have enough count. got=%d\n", count)
	} else {
		fmt.Printf("Not enough. got=%d\n", count)
	}

}
