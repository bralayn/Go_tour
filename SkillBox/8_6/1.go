//Задание 1. Времена года
//
//
//Что нужно сделать
//Пользователь вводит месяц, программа должна вывести, на какое время года (зиму, весну, лето, осень) этот месяц выпадает.
//
//Как группировать:
//
//декабрь, январь, февраль — зима;
//март, апрель, май — весна;
//июнь, июль, август — лето;
//сентябрь, октябрь, ноябрь — осень.
package main

import "fmt"

func main() {

	fmt.Println("Seasons")
	fmt.Println("Input month ")
	var month string
	_, _ = fmt.Scan(&month)

	switch month {
	case "december", "january", "february":
		fmt.Println("Seasons - winter")
	case "march", "april", "may":
		fmt.Println("Seasons - spring")
	case "june", "july", "august":
		fmt.Println("Seasons - summer")
	case "september", "october", "november":
		fmt.Println("Seasons - autumn")

	}
}
