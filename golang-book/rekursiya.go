package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(0))
	fmt.Println(factorial(1))
	fmt.Println(factorial(2))
	fmt.Println(factorial(3))
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
}
