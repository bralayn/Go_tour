package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64

	fmt.Scan(&number)

	if number <= 0 {
		fmt.Printf("число %2.2f не подходит", number)
	} else if number > 10000 {
		fmt.Printf("%e", number)
	} else {
		new_number := math.Pow(number, 2)
		fmt.Printf("%.4f", new_number)

	}
}
