package main

import "fmt"

func main() {

	var numberDay, summaCheck, totalCheck int
	//	discount := 10 / 100

	fmt.Print("Input day of week (1-7): ")
	fmt.Scan(&numberDay)
	fmt.Println()

	fmt.Print("Input check amount: ")
	fmt.Scan(&summaCheck)
	fmt.Println()

	if numberDay < 6 {
		fmt.Println("To pay without a discount: ", summaCheck)
	}

	if summaCheck >= 10000 {
		discount := summaCheck * 10 / 100
		totalCheck = summaCheck - discount

		if numberDay == 6 {
			fmt.Println("Today is saturday! Discount = ", discount)
			fmt.Println(" To pay with a discount: ", totalCheck)
		}
		if numberDay == 7 {
			fmt.Println("Today is sunday! Discount = ", discount)
			fmt.Println(" To pay with a discount: ", totalCheck)
		}
	}
}
