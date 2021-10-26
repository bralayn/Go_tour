package main

import "fmt"

func main() {

	var (
		numberOfman, numberOfbarber int
		needBarber, countMan        int
	)
	days := 30
	hours := 8

	fmt.Print(" Input, how many man lives in the city? :")
	fmt.Scan(&numberOfman)
	fmt.Println()

	fmt.Print("Input, how many barbers work in the city? :")
	fmt.Scan(&numberOfbarber)
	fmt.Println()

	// calculate how many men a 1 barber can serve per month

	fmt.Println("----------- Calculate -------------")
	countMan = hours * days
	needBarber = numberOfbarber * countMan

	if needBarber >= numberOfman {
		fmt.Println("-- there are enough barbers --", needBarber)
	}

	if needBarber < numberOfman {
		fmt.Println("--- there are not enough barbers! ---", needBarber)

	}
}
