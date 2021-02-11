package main

import "fmt"

func main() {
	var n uint
	fmt.Println("Inpun Number n <= 10000 ")
	fmt.Scan(&n)

	n1 := n / 1000
	n2 := n / 100 % 10
	n3 := n / 10 % 10

	switch {
	case n == 10000:
		fmt.Println(1)
	case n1 > 0:
		fmt.Println(n1)
	case n2 > 0:
		fmt.Println(n2)
	case n3 > 0:
		fmt.Println(n3)
	default:
		fmt.Println(n % 10)
	}

	//switch {
	//case n / 1000 > 0:
	//	fmt.Println(n/1000)
	//case n / 100 % 10 > 0:
	//	fmt.Println(n / 100 % 10)
	//case n / 10 % 10 > 0:
	//	fmt.Println(n / 10 %10)
	//default:
	//	fmt.Println(n % 10)
	//}
}
