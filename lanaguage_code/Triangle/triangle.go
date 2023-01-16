package main

import "fmt"

func TriangleCheck(firstSide, secondSide, thirdSide int) string {
	if firstSide <= 0 || secondSide <= 0 || thirdSide <= 0 {
		return "Invalid sides"
	}
	if firstSide+secondSide <= thirdSide || secondSide+thirdSide <= firstSide || thirdSide+firstSide <= secondSide {
		return "Given combination cann't form the triangle "
	}
	if firstSide == secondSide && secondSide == thirdSide {
		return "Equilateral Triangle"
	}
	if firstSide == secondSide || secondSide == thirdSide || thirdSide == firstSide {
		return "Isosceles Triangle"
	}
	return "Scalene Triangle"
}

func main() {
	fmt.Println("Enter side one :")
	var firstSide int
	fmt.Scanf("%d", &firstSide)

	fmt.Println("Enter side two :")
	var secondSide int
	fmt.Scanf("%d", &secondSide)

	fmt.Println("Enter side thirdSide :")
	var thirdSide int
	fmt.Scanf("%d", &thirdSide)

	// How to take multiple input in one go
	// var a int
	// var b int
	// var c int
	// fmt.Scanf("%d %d %d", &a, &b, &c)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)


	result := TriangleCheck(firstSide, secondSide, thirdSide)
	fmt.Println(result)
}
