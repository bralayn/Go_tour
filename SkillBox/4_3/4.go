package main

import "fmt"

func main() {

	var number int

	fmt.Print(" Input experience points : ")
	fmt.Scan(&number)
	fmt.Println()

	if number >= 5000 {
		fmt.Println("Your level = 4")
	} else if number >= 2500 {
		fmt.Println("Your level = 3")
	} else if number >= 1000 {
		fmt.Println("Your level = 2")
	} else {
		fmt.Println("Your level = 1")
	}
}
