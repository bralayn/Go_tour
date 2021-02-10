package main

import "fmt"

func main() {
	var (
		d, m, h int
	)
	fmt.Print("Input d=>")
	fmt.Scan(&d)
	//h = (d * 12 / 360) % 10
	h = (2 * d) / 60
	m = (d * 2) % 60

	fmt.Println("It is", h, "hours", m, "minutes")
}
