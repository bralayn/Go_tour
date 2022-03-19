//Задание 1. Написание простого цикла
//
//
//Что нужно сделать
//Напишите программу, которая будет принимать от пользователя произвольное число и в цикле выводить
//на экран все значения от нуля до указанного числа.
package main

import "fmt"

func main() {
	var number int

	fmt.Print("Input some number=")
	fmt.Scan(&number)
	fmt.Println()

	for i := 0; i < number; i++ {
		fmt.Println("Number [", i, "]=", i)
	}

}
