//Задание 1. Функция, возвращающая результат
//
//Что нужно сделать
//Напишите функцию, которая на вход получает число и возвращает true, если число четное, и false, если нечётное.
//
//Рекомендация
//Программа запрашивает у пользователя или генерирует случайное число, передает в функцию в качестве аргумента
//и выводит в консоль результат её работы.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int

	fmt.Print("Input number =")
	fmt.Scan(&num)
	checkNumber(num)

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(1000)
	fmt.Println("Random number=", x)
	checkNumber(x)

}

func checkNumber(n int) {
	var k bool
	if n%2 == 0 {
		//fmt.Println("True") // 1 - metod
		k = true
		fmt.Println(k) //  2 - metod)
	} else {
		//fmt.Println("False")
		k = false
		fmt.Println(k)
	}
}
