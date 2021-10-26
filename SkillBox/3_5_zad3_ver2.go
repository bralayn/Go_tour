package main

import "fmt"

func main() {
	fmt.Println("Обмен местами значений переменных")
	var c int

	a := 50
	b := 666

	c = a
	a = b
	b = c

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
