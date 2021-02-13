package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Input A=>")
	fmt.Scan(&a)
	fmt.Print("Input B>A =>")
	fmt.Scan(&b)

	sum := 0

	var i int

	for i = a; i <= b; i++ {
		fmt.Print(i)
		sum += i

	}
	fmt.Println()
	fmt.Println(sum)

}
