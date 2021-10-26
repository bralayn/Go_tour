package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("Input a:")
	fmt.Scan(&a)
	fmt.Println()

	fmt.Print("Input b:")
	fmt.Scan(&b)
	fmt.Println()

	if a < b {
		fmt.Println("Minimum = ", a)
	} else if a == b {
		fmt.Println("Number a", a, "= Number b", b)
	} else {
		fmt.Println("Minimum = ", b)
	}
}
