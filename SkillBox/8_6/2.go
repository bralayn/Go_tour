//Задание 2. Дни недели
//
//
//Что нужно сделать
//Пользователь вводит будний день недели в сокращённой форме (пн, вт, ср, чт, пт) и получает развёрнутый список
//всех последующих рабочих дней, включая пятницу.
//
//Рекомендация
//Пример работы программы:
//
//Дни недели.
//Введите будний день недели: пн, вт, ср, чт, пт:
//вт
//вторник
//среда
//четверг
//пятница
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Days of week.")
	fmt.Println("Input weekday of the week: mo, tu, we, th, fr: ")
	var weekDay string
	fmt.Scan(&weekDay)

	switch weekDay {
	case "mo":
		fmt.Println("Monday")
		fallthrough
	case "tu":
		fmt.Println("Tuesday")
		fallthrough
	case "we":
		fmt.Println("Wednesday")
		fallthrough
	case "th":
		fmt.Println("Thursday")
		fallthrough
	case "fr":
		fmt.Println("Friday")

	}
}
