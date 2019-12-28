package main

import (
	"fmt"
)

var pi = 3.14

// y := 15 // raccourci := interdit dans la déclaration d'un fichier !

func main() {
	var age int = 20 // déclaration longue
	// var age int // sans initialisation, age=0
	fmt.Println(age)

	var weight, height int = 80, 190 // déclaration multiple sur une ligne
	fmt.Println(weight, height)

	var name = "Bob" // type string, inféré par la valeur d'initilisation "Bob"
	fmt.Println(name)

	email := "bob@go.org" // := évite raccourci la déclaration (pas besoin de var ni du type)
	fmt.Println(email)

	fmt.Println(pi)
}
