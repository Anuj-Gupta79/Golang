package main

import "fmt"

// make(map[data type of key] data type of value)
func main() {
	mp1 := make(map[int]string)
	mp1[1] = "One"
	mp1[2] = "Two"
	mp1[3] = "Three"
	mp1[5] = "Five"
	mp1[4] = "Four"
	fmt.Println(mp1[1])
	for key, value := range mp1{
		fmt.Println(key, " : ", value);
	}

	mp2 := make(map[string] int)
	mp2["One"] = 1
	mp2["Two"] = 2
	mp2["Three"] = 3
	mp2["A"] = 1
	mp2["Five"] = 5
	fmt.Println(mp2)
	
	// Two check the weather value for a key present or not
	value, isKepPresent := mp2["A"]

	if isKepPresent {
		fmt.Println("The value is present : ", value)
	} else{
		fmt.Println("The value is not present")
	}
}
