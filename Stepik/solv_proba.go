package main

import "fmt"

func main() {
	var num1, num2, n, n1 int

	num1 = 564
	num2 = 564
	//p = 5

	//for ; x < y; i++ {
	//	x=x*(100+p)/100
	//}

	n = num1 / 1000 % 10
	n1 = num2 % 100

	fmt.Println(n, "  ", n1)
}
