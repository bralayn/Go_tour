package main

import (
	"fmt"
)

func main() {
	var (
		superRacecircle int = 4
		shumaher        int = 358
		speedShumaher       = 358
		engineEquip         = 254
		wheelsEquip         = 93
		st_wheelEquip       = 49
		weatherWind         = 21
		weatherRain         = 17
	)

	fmt.Println("===================")
	fmt.Print("Супер гонки.Круг ", superRacecircle, "\n")
	fmt.Println("===================")
	fmt.Print("Шумахер (", shumaher, ")\n")
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Print("Скорость: ", speedShumaher, "\n")
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engineEquip, "\n")
	fmt.Print("Колеса: +", wheelsEquip, "\n")
	fmt.Print("Руль: +", st_wheelEquip, "\n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: -", weatherWind, "\n")
	fmt.Print("Дождь: -", weatherRain, "\n")
}
