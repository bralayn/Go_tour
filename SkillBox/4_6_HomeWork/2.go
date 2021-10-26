//Задание 2. Три числа
//Что нужно сделать
//Напишите программу, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число больше пяти.

package main

import "fmt"

func main() {

	var firstNumber, secondNumber, thirdNumber int

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

	if firstNumber > 5 {
		fmt.Println("Among the numbers entered, there is a number greater than 5 ", "it is: ", firstNumber)
	} else if secondNumber > 5 {
		fmt.Println("Among the numbers entered, there is a number greater than 5 ", "it is: ", secondNumber)
	} else if thirdNumber > 5 {
		fmt.Println("Among the numbers entered, there is a number greater than 5 ", "it is: ", thirdNumber)
	}

	if firstNumber <= 5 && secondNumber <= 5 && thirdNumber <= 5 {
		fmt.Println("There is no number greater than 5 among the entered numbers ")
	}

}
