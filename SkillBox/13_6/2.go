//Задание 2. Функция, принимающая значение по ссылке
//
//Что нужно сделать
//Напишите функцию, которая принимает в качестве аргументов указатели на два типа int и меняет их значения местами.
//
//Рекомендация
//В методе main создайте и присвойте значения двум переменным типа int, выведите значения этих переменных в консоль до вызова функции и после.
package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Input a=")
	fmt.Scan(&a)
	fmt.Print("Input b=")
	fmt.Scan(&b)
	fmt.Println("--- Befor the function call ---")
	fmt.Println("a=", a, "b=", b)

	changePlaces(&a, &b)
	fmt.Println("--- After the function call ---")
	fmt.Println("a=", a, "b=", b)
}

func changePlaces(a, b *int) {
	*a, *b = *b, *a
}
