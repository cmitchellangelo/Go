package main

import "fmt"

func main() {
	//Arrays
	//var fruitArr [2]string

	//Assign values
	//fruitArr[0] = "Apples"
	//fruitArr[1] = "Oranges"

	// Declare and assign
	//fruitArr := [2]string{"Apple", "Orange"}

	//fmt.Println(fruitArr)
	//fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
