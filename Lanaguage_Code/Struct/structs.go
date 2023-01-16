package main

import "fmt"

// Syntax for making structure

// type Struct_Name struct{
// 	...field..
// }

// make a struct vertex
type Vertex struct {
	X int
	Y int
}

func main() {
	// how to declare the struct
	V := Vertex{X : 4, Y : 9} // X=4, Y=9
	fmt.Println(V)

	W := Vertex{11, 12} // X=11, Y=12
	fmt.Println(W)

	Empty := Vertex{} // X=0 , Y=0
	fmt.Println(Empty) 
}
