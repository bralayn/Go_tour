package main

import "fmt"

func main() {

	var i, n, sum, number int

	fmt.Print("Input kol-vo chisel v posledovanosti => ")
	fmt.Scan(&n)
	fmt.Println(n)

	//var number int
	for i = 1; i <= n; i++ {
		fmt.Scan(&number)
		fmt.Print()
		fmt.Print(number)
		if number%8 == 0 && number > 9 && number < 100 {
			//sum = sum + number
			sum += number
		}
	}
	fmt.Println(sum)
}
