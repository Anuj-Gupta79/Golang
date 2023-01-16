package main

import "fmt"

// func HelloName(name string) (greetingMessage string) {
// 	greetingMessage =  "Hello " + name
// 	return 
// }

func HelloName(name string) string {
	return   "Hello " + name
}

func main() {
	greet := HelloName("Anuj")
	fmt.Println(greet)
}

