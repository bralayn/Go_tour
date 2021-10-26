//Напишите программу, которая угадывает число, задуманное пользователем от 1 до 5.
//Программа должна спрашивать, больше или меньше задуманное число (от какого-то другого числа) и на основе ответов
//пользователя определить то, что он загадал.
package main

import "fmt"

func main() {
	var number, temp int

	fmt.Println("------------------------------")
	fmt.Print("think of a number from 1 to 5, After input '0' to continue- ")
	fmt.Scan(&temp)
	fmt.Println()

	fmt.Println("The number is -> ", 3)
	fmt.Print("If number < 3 - input 1, if number > 3, input 2, if number = 3, input 66")
	fmt.Scan(&number)
	fmt.Println()

	if number == 1 {
		fmt.Println(" The number is -> ", 2)
		fmt.Print("If number < 2 - input '0', if number = 2, input '66' ")
		fmt.Scan(&number)
		fmt.Println()
		if number == 0 {
			fmt.Println("The number is -> ", 1)
		} else if number == 66 {
			fmt.Println("The number is -> 2  !!!! ")
			fmt.Println()
		}
	}

	if number == 2 {
		fmt.Println(" The number is -> ", 4)
		fmt.Print("If number > 4 - input '0', if number = 4, input '66' ")
		fmt.Scan(&number)
		fmt.Println()
		if number == 0 {
			fmt.Println("The number is -> ", 5)
		} else if number == 66 {
			fmt.Println("The number is -> 4  !!!! ")
			fmt.Println()
		}

		if number == 66 {
			fmt.Println("The number is -> 3  !!!! ")
			fmt.Println()
		}

	}

}
