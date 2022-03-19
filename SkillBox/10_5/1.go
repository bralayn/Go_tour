//Что нужно сделать
//Напишите программу, вычисляющую ex посредством разложения в ряд Тейлора с заданной пользователем точностью.
//Пользователь вводит значение x и с точностью до какого знака после запятой необходимо вычислить значение функции.
//
//
//Советы и рекомендации
//Получить значение точности (эпсилон) можно, разделив 1 на 10, возведённую в степень,
//равную количеству знаков после запятой.
//
//Для x = 0 и n = 1 вывод должен быть 1
//
//Для x = 1 и n = 3 вывод должен быть 2,7182539682539684
//
//Для x = 1 и n = 5 вывод должен быть 2,7182815255731922
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, n int
	//	var  float64

	fmt.Print("Input x= ")
	fmt.Scan(&x)
	fmt.Println()

	fmt.Print("Input precision, number of decimal places n = ") //точность, колличество знаков после запятой
	fmt.Scan(&n)
	fmt.Println()

	fact := 1
	prevRezult := 0.0
	rezult := 0.0
	epsilon := math.Pow(0.1, float64(n)) // из условия задачи, задается пользователем

	h := 0

	if x != 0 {

		for {
			if h > 0 {
				fact *= h
				//			fmt.Println("fact= ", fact) //считается правильно...проверка для себя
			}

			rezult += math.Pow(float64(x), float64(n)) / float64(fact)
			if rezult-prevRezult < epsilon {
				fmt.Println("For x=", x, " and n=", n, "Result =", rezult)
				break
			}
			h++
			prevRezult = rezult
		}
	} else {
		fmt.Println("Любая функция в нулевой степени равна ЕДЕНИЦЕ = ", 1)
	}
}
