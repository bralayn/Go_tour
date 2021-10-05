package main

import "fmt"

func main() {
	var (
		a int
		b float32
	)
	fmt.Print("Введите любое целое число 'a' => ")
	fmt.Scan(&a)
	fmt.Println()

	fmt.Print("Введите любое число с плавающей точкой 'b' => ")
	fmt.Scan(&b)
	fmt.Println()

	fmt.Println("Число 'a' в квадрате = ", a*a)
	fmt.Print("Число 'b' в квадрате = ", b*b)

}
