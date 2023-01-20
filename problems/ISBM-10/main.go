package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Validation(ISBM string) bool {
	ISBM = strings.ReplaceAll(ISBM, "-", "")

	n := len(ISBM)

	if n != 10 {
		return false
	}

	sum := 0

	if ISBM[n-1] == 'X' {
		sum += 10
	} else {
		sum += int(ISBM[n-1]-'0')
	}

	counter := 2
	for i := n - 2; i >= 0; i-- {
		if !unicode.IsDigit(rune(ISBM[i])) {
			return false
		}

		sum += int(ISBM[i] - '0')*counter
		counter++
	}

	return sum%11==0
}

func main() {

	fmt.Println(Validation("3-598-21508-8"))

}
