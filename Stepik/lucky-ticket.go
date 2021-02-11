package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x1 := n / 100000
	x2 := n / 10000 % 10
	x3 := n / 1000 % 10
	x4 := n / 100 % 10
	x5 := n / 10 % 10
	x6 := n % 10

	sum_First := x1 + x2 + x3
	sum_Last := x4 + x5 + x6

	switch {
	case sum_First == sum_Last:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}

	//fmt.Println(x1, x2, x3, x4, x5, x6)
}
