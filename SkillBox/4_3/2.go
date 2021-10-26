package main

import "fmt"

func main() {
	var a, b, summ int

	fmt.Print(" Input a: ")
	fmt.Scan(&a)
	fmt.Println()

	fmt.Print("Input b: ")
	fmt.Scan(&b)
	fmt.Println()

	fmt.Print("Input summu, a+b= ")
	fmt.Scan(&summ)
	fmt.Println()

	if a+b == summ {
		fmt.Println("All OK")
	} else {
		fmt.Println("All don't right, you a stupid")
	}
}
