package main

import (
	"fmt"

	"training.go/discovergo/data"
)

func main() {
	fmt.Println(data.Name, data.Age)
	// fmt.Println(data.password) // INTERDIT: variable password n'est pas exporté!
}
