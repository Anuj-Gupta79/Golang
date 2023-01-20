package main

import (
	"fmt"
	"strconv"
)

func encoded(code string) string {
	if len(code)<1{
		return code
	}
	count := 0
	res := ""
	for i := 0; i < len(code)-1; i++ {
		if code[i] == code[i+1] {
			count++
		} else {
			count++
			if count == 1 {
				res += string(code[i])
			} else {
				res += strconv.Itoa(count) + string(code[i])
			}
			count = 0
		}
	}

	if count == 0 {
		res += string(code[len(code)-1])
	} else {
		count++
		res += strconv.Itoa(count) + string(code[len(code)-1])
	}
	
	return res
}

func main() {
	fmt.Println(encoded("  hsqq qww  "))
}
