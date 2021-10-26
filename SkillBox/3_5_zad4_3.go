package main

import "fmt"

func main() {
	var (
		totalHeight    int
		seedlingBambuk int
		growBambuk     int
		eatBambuk      int
		day            int
	)
	fmt.Print("Введите высоту саженца бамбука в см: ")
	fmt.Scan(&seedlingBambuk)
	fmt.Println()

	fmt.Print("Введите рост бамбука за день в см:")
	fmt.Scan(&growBambuk)
	fmt.Println()

	fmt.Print("Введите сколько съедают гусеницы за ночь в см:")
	fmt.Scan(&eatBambuk)
	fmt.Println()

	fmt.Print("За какое кол-во дней Вы хотите узнать высоту бамбука?")
	fmt.Scan(&day)
	fmt.Println()

	// Расчитываем рост бамбука
	totalHeight = seedlingBambuk + (growBambuk-eatBambuk)*day
	fmt.Println("Высота бамбука после", day, " дней будет =", totalHeight, "см.")

}
