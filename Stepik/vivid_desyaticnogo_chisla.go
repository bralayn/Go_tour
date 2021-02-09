package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	//f := x / 10    1 sposob
	//k := f % 10

	f := x % 100 // 2 sposob
	k := f / 10

	fmt.Println("It's number", k)
}
