package main

import "fmt"

func main() {
	var n, number, count int

	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		fmt.Println(n)

		if n > number {
			count = 0
			number = n
			count++
		} else if n == number {
			count++
		}

	}
	fmt.Println(count)

}
