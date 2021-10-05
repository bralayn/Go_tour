package main

import "fmt"

func main() {

	allSumma := 6000000
	entrances := 8
	flatsEntrance := 20
	payRent := allSumma / (entrances * flatsEntrance)

	fmt.Println("Сумма указанная в квитанции:", allSumma)
	fmt.Println("Подъездов в доме:", entrances)
	fmt.Println("Квартир в каждом подъезде:", flatsEntrance)
	fmt.Println("----Результат------")
	fmt.Println("Каждая квартира должна заплатить по", payRent, "руб.")
}
