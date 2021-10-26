package main

import "fmt"

func main() {

	var (
		station1, station2, station3, station4        string = "Деповская", "Молодежная", "Солнечная", "Ленина"
		allPeople, inPeople, outPeople, profitability int
		salaryDriver, costGas, taxes, repairCar       int
	)
	allPeople = 0
	costTicket := 20

	fmt.Println("----- Программа симуляции работы маршрутного такси -----")
	fmt.Println()
	fmt.Print("Прибываем на остановку 'Улица ", station1, "'. ", "\n")
	fmt.Println("В салоне пассажиров: ", allPeople)
	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&outPeople)
	fmt.Println()
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&inPeople)
	fmt.Println()
	fmt.Print("Отправляемся с остановки 'Улица ", station1, "'. ", "\n")
	allPeople += inPeople - outPeople
	fmt.Println("В салоне пассажиров: ", allPeople)
	fmt.Println("----- едем -----")
	profitability += inPeople * costTicket

	fmt.Println()
	fmt.Print("Прибываем на остановку 'Улица ", station2, "'. ", "\n")
	fmt.Println("В салоне пассажиров: ", allPeople)
	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&outPeople)
	fmt.Println()
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&inPeople)
	fmt.Println()
	fmt.Print("Отправляемся с остановки 'Улица ", station2, "'. ", "\n")
	allPeople += inPeople - outPeople
	fmt.Println("В салоне пассажиров: ", allPeople)
	fmt.Println("----- едем -----")
	profitability += inPeople * costTicket

	fmt.Println()
	fmt.Print("Прибываем на остановку 'Улица ", station3, "'. ", "\n")
	fmt.Println("В салоне пассажиров: ", allPeople)
	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&outPeople)
	fmt.Println()
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&inPeople)
	fmt.Println()
	fmt.Print("Отправляемся с остановки 'Улица ", station3, "'. ", "\n")
	allPeople += inPeople - outPeople
	fmt.Println("В салоне пассажиров: ", allPeople)
	fmt.Println("----- едем -----")
	profitability += inPeople * costTicket

	fmt.Println()
	fmt.Print("Прибываем на остановку 'Улица ", station4, "'. ", "\n")
	fmt.Println("В салоне пассажиров: ", allPeople)
	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&outPeople)
	fmt.Println()
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&inPeople)
	fmt.Println()
	fmt.Print("Отправляемся с остановки 'Улица ", station4, "'. ", "\n")
	allPeople += inPeople - outPeople
	fmt.Println("В салоне пассажиров: ", allPeople)
	profitability += inPeople * costTicket

	fmt.Println()
	fmt.Println("----- Калькуляция поездки -----")
	// Расчитываем зарплату водителя и затраты
	salaryDriver = profitability / 4
	costGas = profitability / 5
	taxes = profitability / 5
	repairCar = profitability / 5
	myProfit := profitability - salaryDriver - costGas - taxes - repairCar

	fmt.Println()
	fmt.Println("Всего заработали: ", profitability)
	fmt.Println("Зарплата водителя: ", salaryDriver)
	fmt.Println("Расходы на топливо: ", costGas)
	fmt.Println("Налоги: ", taxes)
	fmt.Println("Расходы на ремонт машины: ", repairCar)
	fmt.Println("Итого доход: ", myProfit)

}
