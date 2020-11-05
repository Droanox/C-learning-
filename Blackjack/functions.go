package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/* Structs for Blackjack */

type Card struct {
	Value string
	Suit  string
}

type Player struct {
	Faceup []Card
}

type Dealer struct {
	Faceup   []Card
	Facedown []Card
}

/* Initializing variables */

type Deck []Card

var playerSum int
var dealerSum int

var playerHasAce int8
var dealerHasAce int8
var standCount int8
var hitCount int8
var playerDrawCount int8
var dealerDrawCount int8

var drawing string

func new() (deck Deck) { // creating the deck
	suit := []string{"♥", "♤", "♣", "♢"}

	value := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

	for i := 0; i < len(suit); i++ {
		for j := 0; j < len(value); j++ {
			card := Card{
				Suit:  suit[i],
				Value: value[j],
			}
			deck = append(deck, card)
		}
	}
	return
}

func valueToInt(value string) int {
	if value == "Jack" || value == "Queen" || value == "King" {
		return 10
	} else if value == "Ace" {
		return 11
	} else {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Conversion error")
		}
		return num
	}
}

func shuffleDeck(startDeck Deck) (endDeck Deck) {
	rand.Seed(time.Now().UTC().UnixNano())
	order := rand.Perm(52)
	for i := 0; i < len(startDeck); i++ {
		if i != order[i] {
			startDeck[order[i]], startDeck[i] = startDeck[i], startDeck[order[i]]
		}
	}
	return startDeck
}

func draw(deck Deck) (drawedCard Card) {
	drawedCard = deck[0]
	deck = append(deck[:0], deck[1:]...)
	return drawedCard
}

func (p *Player) PlayerDraw(deck Deck) []Card {
	p.Faceup = append(p.Faceup, draw(deck))
	if p.Faceup[playerDrawCount].Value == "Ace" {
		playerHasAce += 1
	}
	playerDrawCount++
	return p.Faceup
}

func (d *Dealer) DealerDrawFU(deck Deck) []Card {
	d.Faceup = append(d.Faceup, draw(deck))
	if d.Faceup[dealerDrawCount].Value == "Ace" {
		dealerHasAce += 1
	}
	dealerDrawCount++
	return d.Faceup
}

func Finish(ending string) {
	fmt.Println(ending)
	time.Sleep(2 * time.Second)
	os.Exit(0)
}

//func (d *Dealer) DealerDrawFD(deck Deck) []Card {
//	d.Facedown = append(d.Facedown, draw(deck))
//	return d.Facedown
//}
