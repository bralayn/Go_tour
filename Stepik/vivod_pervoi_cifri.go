package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	// x = 486
	//y1 := x % 10
	//y2 := x % 100 / 10
	//y3 := x % 1000 / 100
	//y4 := x % 10000 / 1000

	switch {
	case n/1000 > 0:
		fmt.Println(n / 1000)
	case n/100%10 > 0:
		fmt.Println(n / 100 % 10)
	case n/10%10 > 0:
		fmt.Println(n / 10 % 10)
	default:
		fmt.Println(n % 10)
	}
	//y := n % 100 / 10
	//z := n % 1000 / 100
	//
	//switch {
	//case x != y && x != z && y != z:
	//	fmt.Println("YES")
	//default:
	//	fmt.Println("NO")
	//}
	//
	//fmt.Println(n / 100)

	///fmt.Println(y1, y2, y3, y4)
}
