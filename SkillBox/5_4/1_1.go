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

	switch {
	case x > 0 && y > 0:
		fmt.Println("Point with coordinates x=", x, "y=", y, " is in 1 plane ")
	case x < 0 && y > 0:
		fmt.Println("Point with coordinates x=", x, "y=", y, " is in 2 plane ")
	case x < 0 && y < 0:
		fmt.Println("Point with coordinates x=", x, "y=", y, " is in 3 plane ")
	case x > 0 && y < 0:
		fmt.Println("Point with coordinates x=", x, "y=", y, " is in 4 plane ")
	case x != 0 && y == 0:
		fmt.Println("Point with coordinates x=", x, "y=", y, " lies on the abscissa axis ")
	case x == 0 && y != 0:
		fmt.Println("Point with coordinates x=", x, "y=", y, " lies on the ordinate axis ")
	case x == 0 && y == 0:
		fmt.Println("Point with coordinates x=", x, "y=", y, " lies at the origin  ")
	}

}
