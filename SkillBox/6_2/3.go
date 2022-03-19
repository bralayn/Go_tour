//Задача 3. Возведение в степень
//
//Напишите программу, которая будет возводить в степень число пользователя. Программа должна принимать от пользователя
//число и степень, в которую его нужно возвести. В отрицательную степень возводить не нужно.
//Помните, что любое число в степени ноль равно единице.
package main

import (
	"fmt"
	"math"
)

func main() {
	var number, degree float64

	fmt.Println("--- Raising a number to a degree ---")
	//exitProg := ""

	for i := 0; i < 5; i++ {
		fmt.Print("Input number=")
		fmt.Scan(&number)
		fmt.Println()

		fmt.Print("Input degree>0 , IF degree=00 ->EXIT program, =")
		fmt.Scan(&degree)
		fmt.Println()
		if degree == 00 {
			fmt.Println(" END programm! ")
			return
		}

		rezult := math.Pow(number, degree)
		fmt.Println("-- Rezults= ", rezult, "--")
	}
}
