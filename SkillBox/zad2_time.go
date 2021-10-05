package main

import "fmt"

func main() {
	timeSmena := 432
	timeOrder := 4
	timeCollect := 5
	numberClients := timeSmena / (timeOrder + timeCollect)

	fmt.Println("Эта программа рассчитывает, сколько клиентов успеет обслужить кассир за смену.")
	fmt.Println("Введите длительность смены в минутах:", timeSmena)
	fmt.Println("Сколько минут клиент делает заказ?", timeOrder)
	fmt.Println("Сколько минут кассир собирает заказ?", timeCollect)
	fmt.Println("-----Считаем-----")
	fmt.Println("За смену длинной", timeSmena, "минут кассир успеет обслужить", numberClients, "клиентов.")
}
