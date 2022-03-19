//Задание 3. Расчёт суммы скидки
//
//
//Что нужно сделать
//Напишите программу, которая принимает на вход цену товара и скидку. Посчитайте и верните на экран сумму скидки.
//Скидка должна быть не больше 30% от цены товара и не больше 2000 рублей.
package main

import "fmt"

func main() {
	var summaProduct, discount, allDiscount float64

	maxDiscount := 2000

	fmt.Print("Input product price=")
	fmt.Scan(&summaProduct)
	fmt.Println()

	fmt.Print("Input discount % = ")
	fmt.Scan(&discount)
	fmt.Println()

	righDiscount := discount / 100
	allDiscount = summaProduct * righDiscount

	if allDiscount > float64(maxDiscount) {
		fmt.Println("Cost of product=", summaProduct)
		fmt.Println("You discount= ", maxDiscount)
	} else {
		fmt.Println("Cost of product=", summaProduct)
		fmt.Println("You discount= ", allDiscount)
	}
}
