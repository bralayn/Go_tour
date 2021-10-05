package main

import (
	"fmt"
)

func main() {
	var (
		lap           int = 4
		engine        int
		wheels        int
		steeringWheel int
		wind          int
		rain          int
		speed         int
	)

	fmt.Print("Введите значение двигателя => ")
	fmt.Scan(&engine)
	fmt.Println(engine)
	fmt.Print("Введите значение колеса => ")
	fmt.Scan(&wheels)
	fmt.Println()
	fmt.Print("Введите значение руля => ")
	fmt.Scan(&steeringWheel)
	fmt.Println()
	fmt.Print("Введите значение ветра => ")
	fmt.Scan(&wind)
	fmt.Println()
	fmt.Print("Введите значение дождя => ")
	fmt.Scan(&rain)

	speed = engine + wheels + steeringWheel - wind - rain

	fmt.Println("===================")
	fmt.Print("Супер гонки. Круг ", lap, "\n")
	fmt.Println("===================")
	fmt.Print("Шумахер (", speed, ")\n")
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Print("Скорость: ", speed, "\n")
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engine, "\n")
	fmt.Print("Колеса: +", wheels, "\n")
	fmt.Print("Руль: +", steeringWheel, "\n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: -", wind, "\n")
	fmt.Print("Дождь: -", rain, "\n")
}
