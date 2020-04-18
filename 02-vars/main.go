package main

import (
	"fmt"
	"reflect"
)

func main() {
	//fmt.Println("Variables: ")
	//Strings
	// Intger int8, in16, int32, int64
	//Boolean
	// Float

	// var n bool // default Flase
	// fmt.Printf("%v, %T\n", n, n)

	// var m int // m == 0
	// fmt.Printf("%v, %T\n", m, m)

	// var num int32 = 42
	// fmt.Printf("%v, %T\n", num, num)

	// a := 10             //1010
	// b := 3              //0011
	// fmt.Println(a & b)  // 0010
	// fmt.Println(a | b)  // 1011
	// fmt.Println(a ^ b)  // 1001
	// fmt.Println(a &^ b) //0100

	// c := 8              // 2^3
	// fmt.Println(c << 3) // 2^3 * 2^3 = 2^64
	// fmt.Println(c >> 3) // 2^3 / 2*3 = 2^0
	//var name string = "John"

	var age int
	fmt.Println("How old are you? ")
	fmt.Scan(&age)
	fmt.Printf("You age is: %d\n", age)
	const isCool = true

	// Short must be declare inside function
	//name := "Alex"
	//size := 1.3
	// var size float32 = 2.3
	// //email := "test@gmail.com"

	name, email := "Alex", "alex@gmail.com"

	fmt.Println(name, " is ", age, isCool, email)
	fmt.Println(reflect.TypeOf(name))

	// fmt.Printf("%T\n", name)
	// fmt.Printf("%T", size)
}
