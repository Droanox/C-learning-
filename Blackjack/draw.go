package main

import (
	"fmt"
	"strings"
	"time"
)

func gameDraw(deck Deck, name string, p Player, d Dealer) (newp Player, newd Dealer) {
	hitCount = 2
	standCount = 2
	for {
		fmt.Println()
		fmt.Println("'Hit' or 'Stand' ")
		fmt.Scanln(&drawing)
		switch {
		case (strings.Title(strings.ToLower(drawing)) == "Stand"):
			fmt.Println()
			for {
				if dealerSum == playerSum && dealerSum <= 16 {
					time.Sleep(500 * time.Millisecond)
					fmt.Println(name, ">", p.Faceup)
					fmt.Println("Dealer >", d.DealerDrawFU(deck))
					fmt.Println()
					dealerSum += valueToInt(d.Faceup[standCount].Value)
					standCount++
				}
				if dealerSum < playerSum {
					time.Sleep(500 * time.Millisecond)
					fmt.Println(name, ">", p.Faceup)
					fmt.Println("Dealer >", d.DealerDrawFU(deck))
					fmt.Println()
					dealerSum += valueToInt(d.Faceup[standCount].Value)
					standCount++
				}
				if dealerSum == playerSum && dealerSum > 16 {
					Finish("It's a draw, we're all losers") // Draw
				}
				if dealerSum > playerSum && dealerSum < 21 && standCount != 2 {
					Finish("Dealer wins with skill") // Dealer wins
				}
				if dealerSum > playerSum && dealerSum < 21 && standCount == 2 {
					fmt.Print(name)
					Finish(" can't fucking count, what a dickhead") // Dealer wins
				}
				if dealerSum > 21 && dealerHasAce > 0 { // if the dealer has an ace -10 to not bust
					dealerSum = dealerSum - 10
					dealerHasAce = dealerHasAce - 1
				}
				if dealerSum > 21 {
					Finish("Dealer busted, you win. Must be using gameshark...") // Player wins
				}
				if dealerSum == 21 {
					Finish("Dealer got blackjack wheyyy you're a loser") // Dealer wins
				}
			}
		case (strings.Title(strings.ToLower(drawing)) == "Hit"):
			fmt.Println()
			fmt.Println(name, ">", p.PlayerDraw(deck))
			fmt.Println("Dealer >", d.Faceup)
			fmt.Println()
			playerSum += valueToInt(p.Faceup[hitCount].Value)
			hitCount++
			if playerSum > 21 && playerHasAce > 0 { // if the player has an ace -10 to not bust
				playerSum = playerSum - 10
				playerHasAce = playerHasAce - 1
			}
			if playerSum > 21 {
				fmt.Print(name)
				Finish(" busted what a retard") // Dealer wins
			}
			if playerSum == 21 {
				fmt.Print(name)
				Finish(" managed to pull off a blackjack, lucky fucker") // Player wins
			}
		default:
			fmt.Println("Enter 'Hit' or 'Stand' retard")
			time.Sleep(2 * time.Second)

		}
	}
}
