package main

import (
	"fmt"
	"strings"
)

func CheckValidation(number string) bool{
	str := strings.ReplaceAll(number, " ", "")

	if len(str)<=1 {
		return false
	}

	for characters := range str {
		if characters > '9' || characters<'0' {
			return false
		}
	}

	sum := 0
	for i:=len(str)-2; i>=0; i -=2 {
		if 2*(int(str[i])) > 9 {
			sum += 2*(int(str[i])) - 9 
		} else {
			sum += 2*int(str[i])
		}
	}

	fmt.Println(sum)

	for i:=len(str)-1; i>=0; i-=2 {
		sum += int(str[i])
	}

	if sum%10==0{
		return true
	}

	return false
}

func main() {

	fmt.Println(CheckValidation("8273 1232 7352 0569"))
}
