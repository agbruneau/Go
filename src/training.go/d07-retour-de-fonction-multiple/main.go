package main

import (
	"fmt"
	"strings"
)

func ToLowerStr(name string) (string, int) {
	return strings.ToLower(name), len(name)
}

func main() {
	lowerName, len := ToLowerStr("ALICE")
	fmt.Printf("%s (len=%d)\n", lowerName, len)

	_, len = ToLowerStr("Bob")
	fmt.Printf("bob len=%d\n", len)
}
