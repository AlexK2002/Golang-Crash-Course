package main

import (
	"fmt"
	"strconv"
)

// Person Define Struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting Method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I'm " + strconv.Itoa(p.age)
}

// hasBirhday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" || p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "John", lastName: "Smith", city: "Jerusalem", gender: "M", age: 30}

	//Alternative init
	person2 := Person{"Avi", "Cohen", "Jerusalem", "M", 25}
	person3 := Person{"Tali", "Aviv", "Modiin", "F", 28}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Timtam")
	fmt.Println(person1.greet())

	person3.getMarried("Williams")
	fmt.Println(person3.greet())

}
