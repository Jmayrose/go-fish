package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Competitor - Computer or Human Player

func takeTurn(gotAny int64, player *[]Card, opponent *[]Card) {

	//if opponent contains a card with value gotAny
	if true {
		//take card(s)
	} else {
		// draw from deck
	}
	//if player contains 4 cards with matching values
	if true {
		//resolve matches and remove from hand and increment score
	}
}

func main() {
	playBlackjack()
	type Competitor struct {
		hand         []Card
		score, guess int64
	}

	var deck = makeDeck()
	var player Competitor
	var dealer Competitor

	player.hand = drawCards(player.hand, &deck, 7)
	player.hand = drawCards(dealer.hand, &deck, 7)

	for {
		//TODO Prompt player for card to ask for
		fmt.Println("Enter a card value to ask for")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var guessParse = scanner.Text()
			switch guessParse {
			case "J":
			case "Q":
			case "K":
			case "A":
			default:
				player.guess, _ = strconv.ParseInt(guessParse, 10, 32)

			}

			//TODO Parse 2-9, J, Q, K, A
		}
		takeTurn(player.guess, &player.hand, &dealer.hand)

		//TODO Determine dealerGuess
		dealer.guess = 2
		takeTurn(dealer.guess, &dealer.hand, &player.hand)
	}
}
