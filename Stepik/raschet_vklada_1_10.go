package main

import "fmt"

func main() {
	var i, x, p, y, sum int

	//	sum = 1
	fmt.Scan(&x, &p, &y)
	for i := x; i <= y; i++ {
		x += x * p / 100
		sum++
	}
	fmt.Println(i)

	fmt.Println(sum)
}
