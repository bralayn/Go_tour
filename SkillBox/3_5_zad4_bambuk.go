package main

import "fmt"

func main() {
	var totalHeight int

	seedlingBambuk := 100
	growBambuk := 50
	eatBambuk := 20

	// Расчитываем рост бамбука к середине 3-го дня
	// 2 полных дня + половина дневного роста в третий день (ночь не наступила, гады сожрать ничего не успели)

	totalHeight = seedlingBambuk + (growBambuk-eatBambuk)*2 + growBambuk/2
	fmt.Println("Высота бамбука в середине третьего дня будет =", totalHeight, "см.")

}
