package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := n % 10
	y := n % 100 / 10
	z := n % 1000 / 100

	switch {
	case x != y && x != z && y != z:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}

	//fmt.Println(x, y, z)
}
