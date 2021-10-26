//Что нужно сделать
//Напишите программу, которая запрашивает у пользователя три числа и выводит количество чисел,
//которые больше или равны пяти.

package main

import "fmt"

func main() {

	var firstNumber, secondNumber, thirdNumber, count int

	fmt.Println(" Three numbers.")
	fmt.Print("Input first number: ")
	fmt.Scan(&firstNumber)
	fmt.Println()

	fmt.Print("Input second number: ")
	fmt.Scan(&secondNumber)
	fmt.Println()

	fmt.Print("Input third number: ")
	fmt.Scan(&thirdNumber)
	fmt.Println()

	if firstNumber >= 5 {
		fmt.Println("Among the numbers entered, there is a number >= 5 ")
		count += 1
	}

	if secondNumber >= 5 {
		fmt.Println("Among the numbers entered, there is a number >= 5 ")
		count += 1
	}

	if thirdNumber >= 5 {
		fmt.Println("Among the numbers entered, there is a number >= 5 ")
		count += 1
	}

	if firstNumber < 5 && secondNumber < 5 && thirdNumber < 5 {
		fmt.Println("There is no number >=5 among the entered numbers ")
	}
	fmt.Println("Among the numbers entered", count, " >= 5")
}
