package main

//TODO Eventually move Card to a golang module

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
	fmt.Print("[", card.suit, displayedValue, "]")
}

func printHand(hand []Card) {
	for eachCard := range hand {
		printCard(hand[eachCard])
	}
	println("")
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

func playBlackjack() int {
	var deck = makeDeck()
	var playerHand []Card
	var dealerHand []Card

	playerHand = drawCards(playerHand, &deck, 2)
	dealerHand = drawCards(dealerHand, &deck, 2)

	//TODO Prompt for hit, stand, double down, or split

	if getBlackjackHandValue(playerHand) < 21 {
		for getBlackjackHandValue(dealerHand) < 17 {
			dealerHand = drawCards(dealerHand, &deck, 1)
		}
	}

	printHand(playerHand)
	printHand(dealerHand)
	return resolveBlackjack(playerHand, dealerHand)
}

func resolveBlackjack(player []Card, dealer []Card) int {
	var playerScore = getBlackjackHandValue(player)
	var dealerScore = getBlackjackHandValue(dealer)

	println(playerScore, " ", dealerScore)

	if playerScore > dealerScore {
		println("Player wins")
		return 1
	} else if playerScore == dealerScore {
		println("Push")
		return 0
	} else {
		println("Dealer Wins")
		return -1
	}
}

func getBlackjackHandValue(hand []Card) int {
	var points int = 0
	var workingCardVal int
	for x := 0; x < len(hand); x++ {
		workingCardVal = hand[x].value
		if workingCardVal < 10 {
			if workingCardVal == 1 {
				points += 11
				//TODO Contend with 1/11 split of Aces
				//?Accomplish this by returning array of ints with both possible points
				// points++
			} else {
				points += workingCardVal
			}
		} else {
			points += 10
		}
	}
	return points
}

func main() {
	//TODO Menu to select between different games
	playBlackjack()
}
