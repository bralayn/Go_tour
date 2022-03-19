//Задание 2. Складывание всех чисел кратных 4
//
//Напишите программу, которая в цикле идет от 0 до 40 и складывает только числа кратные 4. Остальные пропускает.
package main

import "fmt"

func main() {

	summaNumbers := 0
	for i := 0; i <= 40; i++ {
		if i%4 != 0 {
			continue
		}
		summaNumbers += i
		fmt.Println(summaNumbers)

	}
	fmt.Println("Summa all numbers multiple of 4 =", summaNumbers)
}
