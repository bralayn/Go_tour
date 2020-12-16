package main

import "fmt"

const (
	Pi = 3.14
	x  = 6.86
)

func main() {
	const World = "Dady"
	fmt.Println("hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	Sum := Pi + x
	fmt.Println(Sum)
}
