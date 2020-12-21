package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Card structure
type Card struct {
	value int
	suit  string
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
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	return deck
}

func drawCards(hand []Card, deck *[]Card, numCards int) []Card {

	newHand := make([]Card, 0)
	for x := 0; x < numCards; x++ {
		newHand = append(newHand, (*deck)[0])
		*deck = (*deck)[1:]
	}
	hand = append(hand, newHand...)
	return hand
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
	fmt.Print("[", card.suit, displayedValue, "]")
}

func printHand(hand []Card) {
	for eachCard := range hand {
		printCard(hand[eachCard])
	}
	fmt.Println("")
}
