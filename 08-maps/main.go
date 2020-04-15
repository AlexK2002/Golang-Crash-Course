package main

import "fmt"

func main() {
	// Define Map
	// emails := make(map[string]string)

	// //Assigne key value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Tom": "tom@gmail.com", "Sharon": "sharon@gmail.com"}
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	//Delete From map
	delete(emails, "Bob")
	fmt.Println(emails)

}
