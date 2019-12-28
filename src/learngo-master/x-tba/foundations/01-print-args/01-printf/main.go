// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// zero-values
	var (
		planet   string
		distance int
		orbital  float64
		hasLife  bool
	)

	// assignment
	planet = "venus"
	distance = 261
	orbital = 224.701
	hasLife = false

	// var (
	// 	planet   = "venus"
	// 	distance = 261
	// 	orbital  = 224.701
	// 	hasLife  = false
	// )

	distance += 5
	distance *= 2
	distance++
	distance--
	orbital++

	// orbital *= 10
	const constFactor = 10
	orbital *= constFactor

	factor := 10
	// orbital *= factor
	orbital *= float64(factor)

	// swiss army knife %v verb
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	// argument indexing - indexes start from 1
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
		planet, distance,
	)

	// why use other verbs than? because: type-safety
	// uncomment to see the warnings:
	//
	// fmt.Printf("Planet: %d\n", planet)
	// fmt.Printf("Distance: %s millions kms\n", distance)
	// fmt.Printf("Orbital Period: %t days\n", orbital)
	// fmt.Printf("Does %v has life? %f\n", planet, hasLife)

	// correct verbs:
	// fmt.Printf("Planet: %s\n", planet)
	// fmt.Printf("Distance: %d millions kms\n", distance)
	// fmt.Printf("Orbital Period: %f days\n", orbital)
	// fmt.Printf("Does %s has life? %t\n", planet, hasLife)

	// precision
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
}
