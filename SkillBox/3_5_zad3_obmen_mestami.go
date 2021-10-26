package main

import "fmt"

func main() {
	fmt.Println("Обмен местами значений переменных")

	a := 50
	b := 666

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	a = 666
	b = 50

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
