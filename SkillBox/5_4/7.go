//Задание 7. Счастливый билет
//
//Что нужно сделать
//Напишите программу, в которую пользователь будет вводить четырёхзначный номер билета, а программа будет выводить,
//является ли он зеркальным, счастливым или обычным билетом.
//
//Советы и рекомендации
//При решении задачи необходимо применить целочисленное деление и остаток от деления. Примеры вывода программы:
//
//1234 -> обычный билет
//
//3425 -> счастливый билет
//
//1221 -> зеркальный билет
package main

import "fmt"

func main() {
	var number, n1, n2, n3, n4 int

	fmt.Print("Input some four-digit number:")
	fmt.Scan(&number)
	fmt.Println()

	// expand the number
	n1 = number / 1000
	n2 = number / 100 % 10
	n3 = number / 10 % 10
	n4 = number % 10

	if n1 < n2 && n2 < n3 && n3 < n4 {
		fmt.Println("-- Regular ticket ---")
	}
	if n1 == n4 && n2 == n3 {
		fmt.Println("--- Mirror ticket ---")
	}
	if n1%2 > 0 && n4%2 > 0 && n2%2 == 0 && n3%2 == 0 {
		fmt.Println("!!! Lucky ticket !!!")
	}

	fmt.Println(n1, n2, n3, n4)

}
