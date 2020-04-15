package main

import "fmt"

func main() {
	// Arrays

	//var fruit [2]string

	// Assign value

	// fruit[0] = "Apple"
	// fruit[1] = "Orange"

	// fmt.Println(fruit)

	// Declare and assign
	fruitArr := [2]string{"Orange", "Apple"}
	fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Orange", "Banna", "Grape"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
