package main

import "fmt"

func main() {
	var (
		Planet, Star, Ranger string
		Money                int
	)
	fmt.Print("Введите название планеты: ")
	fmt.Scan(&Planet)
	fmt.Println()

	fmt.Print("Введите название звездной системы: ")
	fmt.Scan(&Star)
	fmt.Println()

	fmt.Print("Введите имя рейнджера: ")
	fmt.Scan(&Ranger)
	fmt.Println()

	fmt.Print("Введите сумму вознаграждения: ")
	fmt.Scan(&Money)
	fmt.Println()

	fmt.Print("«Как вам,", Ranger, ",известно, мы раса мирная, поэтому на наших военных кораблях используются наёмники с других планет.", "\n")
	fmt.Print("Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты ", Planet, " системы ", Star, ".")

}
