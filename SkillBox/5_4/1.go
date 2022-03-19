//Задание 1. Определение координатной плоскости точки
//
//Что нужно сделать
//Напишите программу, которая поможет пользователю определить, к какой координатной четверти принадлежит точка.
//Пользователь вводит числа X и Y, а программе необходимо вывести, какой координатной четверти принадлежит точка.
//
//Рекомендация
//Используйте логическое умножение.

package main

import "fmt"

func main() {

	var x, y float64

	fmt.Print("Input value X: ")
	fmt.Scan(&x)
	fmt.Println()
	fmt.Print("Input value Y: ")
	fmt.Scan(&y)
	fmt.Println()

	if x > 0 && y > 0 {
		fmt.Println("Point with coordinates x=", x, "y=", y, " is in 1 plane ")
	}
	if x < 0 && y > 0 {
		fmt.Println("Point with coordinates x=", x, "y=", y, " is in 2 plane ")
	}
	if x < 0 && y < 0 {
		fmt.Println("Point with coordinates x=", x, "y=", y, " is in 3 plane ")
	}
	if x > 0 && y < 0 {
		fmt.Println("Point with coordinates x=", x, "y=", y, " is in 4 plane ")
	}

	// determine which axis lies
	if x != 0 && y == 0 {
		fmt.Println("Point with coordinates x=", x, "y=", y, " lies on the abscissa axis ")
	}
	if x == 0 && y != 0 {
		fmt.Println("Point with coordinates x=", x, "y=", y, " lies on the ordinate axis ")
	}

	if x == 0 && y == 0 {
		fmt.Println("Point with coordinates x=", x, "y=", y, " lies at the origin  ")
	}
}
