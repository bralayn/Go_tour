package main

import "fmt"

func main() {
	var i, x1, x2, x3, y3, y2, y1 int

	fmt.Print("Input i=")
	fmt.Scan(&i)

	x1 = i / 100000
	x2 = i / 10000 % 10
	x3 = i / 1000 % 10
	y3 = i / 100 % 10
	y2 = i / 10 % 10
	y1 = i % 10

	fmt.Println("------", i, "-------")

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(y3)
	fmt.Println(y2)
	fmt.Println(y1)

}
