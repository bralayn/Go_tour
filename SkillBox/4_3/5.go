package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Input a: ")
	fmt.Scan(&a)
	fmt.Println()

	fmt.Print("Input b:")
	fmt.Scan(&b)
	fmt.Println()

	if a%b == 0 {
		fmt.Println(" -a- divided by -b- without a remainder ")
	} else {
		fmt.Println("-a- divided by -b- with a remainder ")
	}

	if b%a == 0 {
		fmt.Println(" -b- divided by -a- without a remainder ")
	} else {
		fmt.Println("-b- divided by -a- with a remainder ")
	}
}
