//Задание 3. Проверить, что есть совпадающие числа
//
//Что нужно сделать
//Пользователь при вводе может ошибиться, и многие программы пытаются привлечь внимание к таким ошибкам.
//Например, текстовые редакторы подсвечивают ошибки в словах, электронные таблицы — в формулах.
//Реализуйте программу, которая запрашивает у пользователя три числа и выводит подсказку,
//если два числа или более совпадают.
//
//Рекомендация
//Используйте логическое сложение.

package main

import "fmt"

func main() {

	var x, y, z float64

	fmt.Print("Input X: ")
	fmt.Scan(&x)

	fmt.Print("Input Y: ")
	fmt.Scan(&y)

	fmt.Print("Input Z: ")
	fmt.Scan(&z)

	if x == y || x == z || y == z {
		fmt.Println("There are the same numbers among the entered numbers ")
	} else {
		fmt.Println("There aren't the same numbers among the entered numbers ")
	}

}
