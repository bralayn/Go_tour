package main

import "fmt"

func main() {

	var number int

	fmt.Print("Input some number:")
	fmt.Scan(&number)
	fmt.Println()

	if number < 0 {
		rezult := -number
		fmt.Println("Modulus of entered number:", rezult)
	}
	if number >= 0 {
		rezult := number
		fmt.Println("Modulus of entered number", rezult)
	}
}
