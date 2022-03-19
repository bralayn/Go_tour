//Задание 1. Функция, принимающая аргументы
//
//Что нужно сделать
//Напишите функцию, которая принимает в качестве аргументов два числа типа int, вычисляет сумму чётных чисел заданного
//диапазона и выводит результат в консоль.
//
//Рекомендация
//Если первое число, переданное в качестве аргумента, будет больше второго, просто поменяйте их местами.
package main

import "fmt"

func sumaNumbers(x1, x2 int) int {

	sum := 0
	if x1 > x2 {
		x2, x1 = x1, x2
	}
	for i := x1; i <= x2; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	var x1, x2 int
	fmt.Print("Input x1= ")
	fmt.Scan(&x1)
	fmt.Print("Input x2= ")
	fmt.Scan(&x2)

	if x1 > x2 {
		x2, x1 = x1, x2
	}

	res := sumaNumbers(x1, x2)

	fmt.Printf("x1= %d\nx2= %d\nsumma= %d\n", x1, x2, res)
}
