package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func splitFloat(sum float32) (a, b float32) {
	a = sum * 5 / 23
	b = sum + a
	return
}
func main() {
	fmt.Println(split(17))
	fmt.Println(splitFloat(20.25))
}
