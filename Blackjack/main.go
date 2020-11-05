package main

import (
	"fmt"
)

func main() {
	/* Creating a name */
	var name string
	fmt.Print("Enter your username: ")
	fmt.Scanln(&name)

	/* Creating + shuffling the deck */
	deck := new()
	shuffleDeck(deck)

	/* Dealing cards out to the Dealer and Player */
	var p Player
	var d Dealer

	p, d = gameDeal(deck, name, p, d)

	/* Player and Dealer drawing cards */
	gameDraw(deck, name, p, d)
}
