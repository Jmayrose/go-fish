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

func printHand(hand []Card) {
	for eachCard := range hand {
		printCard(hand[eachCard])
	}
}

func makeDeck() []Card {
	var deck []Card
	for x := 0; x < 52; x++ {
		var newCard Card
		newCard.value = x%13 + 1
		if x < 13 {
			newCard.suit = "♠"
		} else if x < 26 {
			newCard.suit = "♦"
		} else if x < 39 {
			newCard.suit = "♥"
		} else {
			newCard.suit = "♣"
		}
		deck = append(deck, newCard)
	}
	return deck
}

func drawCards(deck []Card, numCards int) []Card {
	newHand := make([]Card, 0)
	for x := 0; x < numCards; x++ {
		newHand = append(newHand, deck[0])
		deck = deck[1:]
	}
	return newHand
}

//Shuffle passed array
func Shuffle(vals []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(vals), func(i, j int) { vals[i], vals[j] = vals[j], vals[i] })
}

func main() {

	var deck = makeDeck()
	Shuffle(deck)

	println(len(deck))
	var hand1 = drawCards(deck, 5)
	//TODO remove cards from deck when drawing
	//? pass by reference
	println(len(deck))

	// println(hand1)
	printHand(hand1)
}
