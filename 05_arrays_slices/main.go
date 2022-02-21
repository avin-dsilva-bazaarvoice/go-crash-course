package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])

	// Declare and assign
	fruitArr1 := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArr1[0])
	fmt.Println(fruitArr1[1])

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
