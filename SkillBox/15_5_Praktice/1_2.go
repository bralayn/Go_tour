//Задание 1. Подсчёт чётных и нечётных чисел в массиве
//
//Что нужно сделать
//
//Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество чётных и нечётных чисел.
//Для ввода и подсчёта используйте разные циклы.
//
//Что оценивается
//
//Для введённых чисел 1, 1, 1, 2, 2, 2, 3, 3, 3, 4 программа должна вывести: чётных — 4, нечётных — 6.
package main

import "fmt"

func main() {
	var (
		arr                     [10]int
		evenNumbers, oddNumbers int
	)
	for i := 0; i < 10; i++ {
		fmt.Printf("Input number arr[%v] ", i+1)
		fmt.Scan(&arr[i])
	}

	for _, num := range arr {
		if num%2 == 0 {
			evenNumbers += 1
		} else {
			oddNumbers += 1
		}
	}
	fmt.Println(arr)
	fmt.Println("Number of even = ", evenNumbers)
	fmt.Println("Number of odd = ", oddNumbers)
}
