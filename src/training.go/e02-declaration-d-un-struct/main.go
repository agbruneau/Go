package main

import (
	"fmt"

	"training.go/discovergo/player"
)

func main() {
	var p1 player.Player
	p1.Name = "Bob"
	p1.Age = 10

	fmt.Printf("Player 1: %v\n", p1)
	fmt.Printf("p1.Name=%v, p1.Age=%v\n", p1.Name, p1.Age)

	a := player.Avatar{"http://avatar.jpg"}
	fmt.Printf("avatar : %v\n", a)

	p3 := player.Player{
		Name: "John",
		// Age:  25,
		Avatar: player.Avatar{
			Url: "http://url.com",
		},
	}
	fmt.Printf("Player 3: %v\n", p3)

	p4 := player.New("Bobette")
	fmt.Printf("Player 4: %v\n", p4)
}
