package main

import "fmt"

func main() {
	var n, c, d, i int

	fmt.Print("Input n= ")
	fmt.Scan(&n)
	fmt.Print("Input c= ")
	fmt.Scan(&c)
	fmt.Print("Input d= ")
	fmt.Scan(&d)

	for i = 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}

}
