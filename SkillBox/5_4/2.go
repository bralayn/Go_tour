//Задание 2. Проверить, что одно из чисел — положительное
//
//Что нужно сделать
//Проверка пользовательского ввода на различные ограничения является частой задачей. Попросите пользователя
//ввести три числа и проверьте, что хотя бы одно является положительным. Результат проверки необходимо сообщить пользователю.
//
//Рекомендация
//Используйте логическое сложение.

package main

import "fmt"

func main() {

	var x, y, z float64

	fmt.Print("Input X: ")
	fmt.Scan(&x)

	fmt.Print("Input Y: ")
	fmt.Scan(&y)

	fmt.Print("Input Z: ")
	fmt.Scan(&z)

	if x > 0 || y > 0 || z > 0 {
		fmt.Println("There is a positive number among the entered numbers ")
	} else {
		fmt.Println("There isn't a positive number among the entered numbers")
	}

}
