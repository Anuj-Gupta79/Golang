package main

import (
	"errors"
	"fmt"
)

var err = errors.New("ERROR: Number should be greater than zero")
type Classification int

const (
	ClassificationDeficient Classification = 0
	ClassificationPerfect   Classification = 1
	ClassificationAbundant  Classification = 2
	ClassificationError Classification = -1
)

func Type(n int64) (Classification, error) {
	if n<=0 {
		return ClassificationError, err;
	}


	var sum int64

	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	if sum == n {
		return ClassificationPerfect, nil
	}
	if sum > n {
		return ClassificationAbundant, nil
	}

	return ClassificationDeficient, nil
}

func main() {
	var n int64
	fmt.Printf("Enter the number: ")
	fmt.Scanf("%d", &n)

	fmt.Println(Type(n))
}

// func Class(n int64) string {

// 	var sum int64
// 	sum = 0

// 	var i int64
// 	for i = 1; i < n; i++ {
// 		if n%i == 0 {
// 			sum += i
// 		}
// 	}

// 	if sum == n {
// 		return "Perfect Number"
// 	}
// 	if sum > n {
// 		return "Abundant NUmber"
// 	}

// 	return "Deficient NUmber"
// }

// func main() {
// 	var n int64
// 	fmt.Printf("Enter the number: ")
// 	fmt.Scanf("%d", &n)

// 	fmt.Println(Classification(n))
// }
