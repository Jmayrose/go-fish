package main

import (
	"bufio"
	"fmt"
	"os"
)

func playHand(hand *[]Card, deck *[]Card) {
	var canSplit bool = (*hand)[0].suit == (*hand)[1].suit

	//TODO Keep prompting until bust or stand
	for {
		fmt.Println("[1] Hit")   //hit
		fmt.Println("[2] Stand") //stand
		if canSplit {
			fmt.Println("[3] Split") //split
		}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {

			switch scanner.Text() {
			case "1":
				//Hit
				fmt.Println("Hit")
			case "2":
				//Stand
				fmt.Println("Stand")
			case "3":
				if canSplit {
					//Split
					fmt.Println("Split")
					//TODO Play new hands recursively
				}
			}
		}
	}
}

func resolveHand(player []Card, dealer []Card) int {
	var playerScore = getHandValue(player)
	var dealerScore = getHandValue(dealer)

	printHand(player)
	printHand(dealer)
	fmt.Println(playerScore, " ", dealerScore)

	if dealerScore > 21 {
		fmt.Println("Dealer Busts, Player wins")
		return 1
	} else if playerScore > dealerScore {
		fmt.Println("Player Beats Dealer, Player wins")
		return 1
	} else if playerScore == dealerScore {
		fmt.Println("Push")
		return 0
	} else {
		fmt.Println("Dealer Wins")
		return -1
	}
}

func getHandValue(hand []Card) int {
	var points int = 0
	var workingCardVal int
	for x := 0; x < len(hand); x++ {
		workingCardVal = hand[x].value
		if workingCardVal < 10 {
			if workingCardVal == 1 {
				points += 11
				//TODO Contend with 1/11 split of Aces
				//?Accomplish this by returning array of ints with both possible points??
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

func playBlackjack() int {
	var deck = makeDeck()
	var playerHand []Card
	var dealerHand []Card

	playerHand = drawCards(playerHand, &deck, 2)
	dealerHand = drawCards(dealerHand, &deck, 2)

	playHand(&playerHand, &deck)

	if getHandValue(playerHand) > 21 {
		fmt.Println("Player Busts, Dealer Wins")
		return -1
	}
	for getHandValue(dealerHand) < 17 {
		dealerHand = drawCards(dealerHand, &deck, 1)
	}

	return resolveHand(playerHand, dealerHand)
}
