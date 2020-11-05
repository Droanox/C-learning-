package main

import (
	"fmt"
	"time"
)

func gameDeal(deck Deck, name string, p Player, d Dealer) (newp Player, newd Dealer) {
	fmt.Println(name, ">", p.PlayerDraw(deck))
	time.Sleep(500 * time.Millisecond) // time intervals for "drawing effect"
	fmt.Println("Dealer >", d.DealerDrawFU(deck))
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println(name, ">", p.PlayerDraw(deck))
	time.Sleep(500 * time.Millisecond)
	playerSum += (valueToInt(p.Faceup[0].Value) + valueToInt(p.Faceup[1].Value)) // total for players cards for now and drawing
	if playerSum == 21 {
		fmt.Println()
		fmt.Print(name)
		Finish(" got Blackjack, ez win") // Player wins outright
	}
	fmt.Println("Dealer >", d.DealerDrawFU(deck))
	time.Sleep(500 * time.Millisecond)
	dealerSum += (valueToInt(d.Faceup[0].Value) + valueToInt(d.Faceup[1].Value)) // total for dealers cards for now and drawing
	if dealerSum == 21 {
		fmt.Println()
		Finish("Dealer got Blackjack, ez loss") // Dealer wins
	}
	return p, d
}
