package main

import "fmt"

func main() {
	var (
		productFirst, productSecond, productThird, totalCost float64
	)
	discountPecent := 0.1

	fmt.Print("Input price first product: ")
	fmt.Scan(&productFirst)
	fmt.Println()

	fmt.Print("Input price second product: ")
	fmt.Scan(&productSecond)
	fmt.Println()

	fmt.Print("Input price third product: ")
	fmt.Scan(&productThird)
	fmt.Println()

	// calculate the cost of goods

	totalCost = productFirst + productSecond + productThird

	fmt.Println("------------------------------------")
	// check the amount of the check
	if totalCost > 10000 {
		fmt.Println("Cost of all goods=", totalCost)
		fmt.Println("Discount 10% :", totalCost*discountPecent)
		totalCost = totalCost - (totalCost * discountPecent)
		fmt.Println("Discount value of goods=", totalCost)
	}
	if totalCost <= 10000 {
		fmt.Println("Cost of goods=", totalCost)
	}
}
