package main

import "fmt"

func main() {
	var yourCard1 string
	var yourCard2 string
	var dealerCard string

	fmt.Scanln(&yourCard1)
	fmt.Scanln(&yourCard2)
	fmt.Scanln(&dealerCard)

	if yourCard1 == "ace" && yourCard2 == "ace" {
		fmt.Println("P")
	} else if isBlackjack(yourCard1, yourCard2) == true && parseCard(dealerCard) < 10 {
		fmt.Println("W")
	} else {
		fmt.Println("S")
	}

	if (parseCard(yourCard1) + parseCard(yourCard2)) >= 17 {
		fmt.Println("S")
	} else if (parseCard(yourCard1) + parseCard(yourCard2)) <= 11 {
		fmt.Println("H")
	} else if (parseCard(yourCard1)+parseCard(yourCard2)) >= 12 && (parseCard(yourCard1)+parseCard(yourCard2)) <= 16 {
		if parseCard(dealerCard) >= 7 {
			fmt.Println("H")
		} else {
			fmt.Println("S")
		}
	}

}

func parseCard(cardName string) (ans int) {
	switch cardName {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

func isBlackjack(card1, card2 string) bool {
	return (parseCard(card1)+parseCard(card2) == 21)
}