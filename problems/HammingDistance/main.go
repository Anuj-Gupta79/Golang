package main

import "fmt"

func hamming(s1, s2 string) int {
	count := 0
	n := len([]rune(s1))
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}

func main() {
	var s1 string
	var s2 string
	fmt.Scanf(s1, s2)
	fmt.Print(hamming(s1, s2))
}
