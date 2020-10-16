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

func determineSuit(index int) string {
	if index < 13 {
		return "♠"
	} else if index < 26 {
		return "♦"
	} else if index < 39 {
		return "♥"
	} else {
		return "♣"
	}
}

//Shuffles passed array
func Shuffle(vals []Card) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func main() {

	var deck [52]Card
	for x := 0; x < len(deck); x++ {
		deck[x].value = x%13 + 1
		deck[x].suit = determineSuit(x)
	}

	Shuffle(deck)
	for eachCard := range deck {
		printCard(deck[eachCard])
	}
}
