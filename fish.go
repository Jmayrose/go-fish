package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Card object
type Card struct {
	value int
	suit  string
	//TODO Maybe change string to char?
	//? does char exist in golang?
}

func printCard(card Card) {
	var displayedValue = fmt.Sprint(card.value)
	switch card.value {
	case 1:
		displayedValue = "A"
	case 11:
		displayedValue = "J"
	case 12:
		displayedValue = "Q"
	case 13:
		displayedValue = "K"
	}
	fmt.Println(card.suit, displayedValue)
}

func makeDeck() [52]Card {
	var deck [52]Card
	for x := 0; x < len(deck); x++ {
		deck[x].value = x%13 + 1
		if x < 13 {
			deck[x].suit = "♠"
		} else if x < 26 {
			deck[x].suit = "♦"
		} else if x < 39 {
			deck[x].suit = "♥"
		} else {
			deck[x].suit = "♣"
		}
	}
	return deck
}

//Shuffle passed array
func Shuffle(vals []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(vals), func(i, j int) { vals[i], vals[j] = vals[j], vals[i] })
}

func main() {

	var deck = makeDeck()

	for eachCard := range deck {
		printCard(deck[eachCard])
	}

	println("\n")

	//? Have to pass slices to function?
	//? what is a golang slice
	Shuffle(deck[:])
	for eachCard := range deck {
		printCard(deck[eachCard])
	}

}
