package main

import (
	"fmt"
	"strings"
)

func IsIsograms(s string) bool {
	for index, character := range s {
		if string(character) == " " || string(character) == "-" {
			continue
		}
		if strings.ContainsRune(s[index+1:], character) {
			return false
		}
	}
	return true;
}

func main() {
	var s string
	fmt.Scanf(s)
	fmt.Print(IsIsograms(s))
}
