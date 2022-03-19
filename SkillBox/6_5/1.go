//Задание 1. Вывод кратных чисел
//
//Напишите программу, которая с помощью цикла выводит все числа кратные 2 в диапазоне от 0 до 10.
//Используйте continue для пропуска не кратных чисел.
package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {

		if i%2 != 0 {
			continue
		}
		fmt.Println("Number multiple of 2 -> ", i)
	}
}
