package main

import "fmt"

func main() {
	var totalHeight, day int

	seedlingBambuk := 100
	growBambuk := 50
	eatBambuk := 20

	totalHeight = seedlingBambuk
l1:
	if totalHeight <= 300 {
		totalHeight += growBambuk
		day += 1
		//	fmt.Println(totalHeight, day)
		if totalHeight < 300 {
			totalHeight -= eatBambuk
			//		fmt.Println(totalHeight, day)
		}
	}
	if totalHeight < 300 {
		goto l1
	}
	fmt.Println("----------------- Результат работы программы -------------------")
	fmt.Println("Срубить и продать бамбук можно будет через", day, "полных дней.")
}
