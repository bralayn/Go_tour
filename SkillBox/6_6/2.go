//Задание 2. Сумма двух чисел по единице
//
//
//Что нужно сделать
//Напишите программу, которая запрашивает у пользователя два числа и складывает их в цикле следующим образом:
//берёт первое число и прибавляет к нему по единице, пока не достигнет суммы двух чисел.
package main

import "fmt"

func main() {

	var a, b, sumBothnumber int

	fmt.Print("Input a=")
	fmt.Scan(&a)
	fmt.Println()

	fmt.Print("Input b=")
	fmt.Scan(&b)
	fmt.Println()

	sumBothnumber = a

	for i := a; i < (a + b); i++ {
		sumBothnumber += 1
		fmt.Println(sumBothnumber)
	}

	fmt.Println("Summa a+b=", sumBothnumber)
}
