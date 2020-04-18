package main

import (
	"fmt"
)

func main() {
	languages := [...]string{"English", "French", "Urdu", "Farsi", "Hungarian", "Hebrew"}

	for i := range languages {
		fmt.Printf("%d) %v\n", i+1, languages[i])
	}
	fmt.Println("Select a number: ")
	var c int
	fmt.Scan(&c)

	switch c {
	case 1:
		fmt.Println("Selection: English")
	case 2:
		fmt.Println("Selection: French")
	case 3:
		fmt.Println("Selection: Urdu")
	case 4:
		fmt.Println("Selection: Farsi")
	case 5:
		fmt.Println("Selection: Hungarian")
	case 6:
		fmt.Println("Selection: Hebrew")
	default:
		fmt.Println("Wrong Number selection")
	}

}
