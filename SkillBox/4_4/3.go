package main

import "fmt"

func main() {

	var year int

	fmt.Print("Input year:")
	fmt.Scan(&year)
	fmt.Println()

	if year%4 == 0 {
		fmt.Println("Year", year, " is visokosni - 366 days")
	} else {
		fmt.Println("Year", year, " isn't  visokosni - 365 days")
	}

}
