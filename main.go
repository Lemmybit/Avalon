package main

import "fmt"

func main() {

	player := newChar("Joe", SAMURAI, HUMAN)

	fmt.Println("Welcome you Avalon!")
	fmt.Println("You are a " + roletoText(player.role) + "!")
}
