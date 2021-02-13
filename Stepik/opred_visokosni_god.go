package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x1 := n % 400
	x2 := n % 4
	x3 := n % 100

	if x1 == 0 {
		fmt.Println("YES")
	} else if x2 == 0 && x3 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	//switch {
	//
	//case x1 == 0: fmt.Println("YES")
	//case x2 == 0 && x3 != 0: fmt.Println("YES")   //  2 sposob
	//default:
	//	fmt.Println("NO")
	//}

	fmt.Println(x1, x2, x3)
	// fmt.Println(1000 % 400, 1000 % 4)
}
