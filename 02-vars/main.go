package main

import "fmt"

func main() {
	//fmt.Println("Variables: ")
	//Strings
	// Intger
	//Boolean
	// Float

	//var name string = "Alex"
	var age = 30
	const isCool = true

	// SHort must be declare inside function
	//name := "Alex"
	//size := 1.3
	var size float32 = 2.3
	//email := "test@gmail.com"

	name, email := "Alex", "alex@gmail.com"

	fmt.Println(name, age, isCool, email)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T", size)
}
