package main

import (
	"fmt"
	"time"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Alex"))
	fmt.Println(time.Second)
	fmt.Printf("%T\n", time.Second)
	time.Sleep(2 * time.Second)
	fmt.Println(getSum(1, 1))

}
