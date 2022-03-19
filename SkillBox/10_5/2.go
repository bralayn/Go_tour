//Задание 2. Проблемы округления процентов
//
//Что нужно сделать
//В связи с особенностями хранения дробных чисел они не очень подходят для хранения денежных значений (может потеряться
//значимая часть при переполнении, да и дробная часть больше двух знаков не нужна). Но мы попробуем решить задачу
//начисления процентов, используя именно их.
//
//Пусть пользователь вводит сумму, которую он кладёт в банк, ежемесячный процент капитализации и количество лет,
//в течение которых будет открыт вклад.
//Необходимо ежемесячно пересчитывать сумму, округляя её до целого количества копеек в меньшую сторону.
//Например, если после начисления процентов остаток включает 35,78 копейки, то оставляем только 35 копеек,
//а дробную часть отбрасываем.
//По окончании работы программы необходимо вывести, какую сумму получит вкладчик на руки и какая сумма будет зачислена
//в пользу банка за счёт округления копеек.
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		percent, years, allMonth int
		summa, drPercent         float64
	)

	fmt.Print("Input how money = ")
	fmt.Scan(&summa)
	fmt.Println()

	fmt.Print("Input capitalization percentage per month  % = ")
	fmt.Scan(&percent)
	drPercent = float64(percent) / 100
	fmt.Println()

	fmt.Print("Input for how many years = ")
	fmt.Scan(&years)
	allMonth = years * 12
	fmt.Println()

	tempPercent := 0.0
	tempSumma := 0.0

	for i := 1; i <= allMonth; i++ {
		tempPercent = summa * drPercent
		realPercent := math.Floor(tempPercent*100) / 100
		tempSumma += tempPercent - realPercent

		summa += realPercent
	}
	fmt.Println("Summa vklada =", math.Round(summa*100)/100)
	fmt.Println("The bank stole money =", math.Round(tempSumma*10000)/10000)

}
