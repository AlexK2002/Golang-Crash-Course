package main

import "fmt"

func main() {
	ids := []int{23, 43, 545, 56, 45, 12}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	//Range with Maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Tom": "tom@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s \n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: ", k)
	}

}
