package main

import (
	"fmt"
	// "sort"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := ""
	for i := 1; i <= numStarsPerLine; i++ {
		border += "*"
	}

	return fmt.Sprintf("%v \n%v \n%v", border, welcomeMsg, border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	oldMsg = strings.TrimLeft(oldMsg, " ")
	oldMsg = strings.TrimRight(oldMsg, " ")
	// sort.Strings(word)
	return oldMsg
}

func main() {
	fmt.Println(WelcomeMessage("Judy"))
	fmt.Println(AddBorder("welcome!", 10))
	fmt.Println(CleanupMessage("***************************    BUY NOW, SAVE 10%   ***************************"))
}
