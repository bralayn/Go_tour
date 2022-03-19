//Задача 2. Сложение произвольных чисел
//
//Напишите программу, которая в цикле складывает 4 произвольных числа. Программа должна, в теле цикла, спрашивать
//у пользователя новое число, прибавлять его к уже полученной сумме и после 4-х итераций — выводить сумму на экран.
package main

import "fmt"

func main() {
	var number int
	summaNumbers := 0

	for i := 0; i < 4; i++ {
		fmt.Print(" Input some number=")
		fmt.Scan(&number)
		fmt.Println()
		summaNumbers += number
		fmt.Println("---", summaNumbers)
	}
	fmt.Println("Summa all entered number=", summaNumbers)
}
