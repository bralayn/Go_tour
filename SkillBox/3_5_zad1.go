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
	fmt.Print("Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты ", Planet, " системы ", Star, ".", "\n")
	fmt.Println()

	fmt.Println("Но случилась неприятность: в связи с большими потерями в последнее время престиж профессии сильно упал, и теперь не так-то")
	fmt.Print("просто завербовать призывников. Главный комиссар планеты ", Planet, " впрочем, отлично справлялся, но недавно его наградили", "\n")
	fmt.Println("орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!")
	fmt.Println()

	fmt.Print(Ranger, ",вы должны прибыть на планету ", Planet, " системы ", Star, " и помочь выполнить план призыва. Мы готовы выплатить вам", "\n")
	fmt.Print("премию в ", Money, " кредитов за эту маленькую услугу»")

}