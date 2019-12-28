package main

import (
	"fmt"
)

func UpdateVal(val string) {
	val = "value"
}

func UpdatePtr(ptr *string) {
	*ptr = "pointer"
}

func main() {
	i := 1
	var p *int = &i

	fmt.Printf("i=%v\n", i)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("*p=%v\n", *p)
	fmt.Println("---------------")

	s := "Bob"
	sPtr := &s
	s2 := *sPtr
	fmt.Println("String pointer")
	fmt.Printf("*s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("---------------")

	*sPtr = "Alice"
	fmt.Println("Dereference and update")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("---------------")

	UpdateVal(s)
	fmt.Println("Func UpdateVal()")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("---------------")

	UpdatePtr(&s)
	UpdatePtr(sPtr)
	fmt.Println("Func UpdatePtr()")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("---------------")
}
