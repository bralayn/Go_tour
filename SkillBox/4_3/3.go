package main

import "fmt"

func main() {
	var number int

	fmt.Print("Input number: ")
	fmt.Scan(&number)
	fmt.Println()

	temp := number % 2
	if temp == 0 {
		fmt.Println("Even number !")
	} else {
		fmt.Println("Odd number")
	}

}
