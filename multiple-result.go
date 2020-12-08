package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func anyMyProbe(x, y int) (int, int) {
	return x, y
}

func main() {
	a, b := swap("hello", "world")
	c, d := anyMyProbe(45, 55)
	fmt.Println(a, b)
	fmt.Println(b, a)
	fmt.Println(c, d)
}
