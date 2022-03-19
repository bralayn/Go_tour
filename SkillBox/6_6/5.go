//Задание 5 (по желанию). Задача на постепенное наполнение корзин разной ёмкости (if и continue и break)
//
//Что нужно сделать
//Представьте, что у вас есть три корзины разной ёмкости. Пользователю предлагается ввести, какое количество яблок
//помещается в каждую корзину. После этого программа должна заполнить все корзины по очереди, учитывая,
//какие корзины уже заполнены, строго соблюдая очерёдность заполнения и добавляя по одному яблоку в каждой итерации.
//
//Советы и рекомендации
//Используйте if и continue и break.
//
//Пример: пользователь решил, что у корзин будет ёмкость на шесть, четыре и девять яблок.
//Программа должна заполнить корзину 1 и в следующей итерации перейти к корзине 2,
//далее в следующей итерации перейти к корзине 3. Если очередная корзина уже заполнена, программа должна
//переходить к следующей по очереди, и так по кругу, пока не заполнит все.
package main

import "fmt"

func main() {

	var a, b, c, basket1, basket2, basket3 int

	fmt.Print("Input basket capacity N1  =")
	fmt.Scan(&basket1)
	fmt.Println()

	fmt.Print("Input basket capacity N2  =")
	fmt.Scan(&basket2)
	fmt.Println()

	fmt.Print("Input basket capacity N3  =")
	fmt.Scan(&basket3)
	fmt.Println()

	allApple := basket1 + basket2 + basket3

	for i := 0; i < allApple; i++ {

		if a != basket1 {
			a += 1
			fmt.Println("In basket N1 = ", a, "apple")
			continue
		}

		if b != basket2 {
			b += 1
			fmt.Println("In basket N2 = ", b, "apple")
			continue
		}

		if c != basket3 {
			c += 1
			fmt.Println("In basket N3 = ", c, "apple")
			continue
		}

	}
	fmt.Println("a===", basket1)
	fmt.Println("b===", basket2)
	fmt.Println("c===", basket3)
	fmt.Println("All apples= ", allApple)
}
