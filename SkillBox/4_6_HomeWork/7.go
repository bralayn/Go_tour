//Задание 7 (по желанию). Прогрессивный налог
//
//Что нужно сделать
//Напишите программу, которая спрашивает у пользователя его доход и рассчитывает сумму
//налога для него по следующим правилам:
//
//налоговая ставка 30% на доход выше 50 000 рублей означает, что 30% уплачивается не со всей суммы,
//а лишь с той её части, которая превосходит 50 000 рублей;
//ставка 20% на доход от 10 000 до 50 000 рублей обязывает уплатить 20% лишь с той части суммы,
//которая превосходит 10 000 рублей, но не превосходит 50 000 рублей.

package main

import "fmt"

func main() {

	var moneyIncom, progressiveTax int

	fmt.Print("Input your income: ")
	fmt.Scan(&moneyIncom)
	fmt.Println()

	thirteenPercent := 10000 * 13 / 100
	twentyPercent := 40000 * 20 / 100
	thirtyPercent := (moneyIncom - 50000) * 30 / 100

	if moneyIncom > 50000 {
		progressiveTax = thirtyPercent + twentyPercent + thirteenPercent
		fmt.Println("From incom: ", moneyIncom, " rub  have to pay: ", progressiveTax, " rub")
	} else if moneyIncom > 10000 && moneyIncom <= 50000 {
		progressiveTax = (moneyIncom-10000)*20/100 + thirteenPercent
		fmt.Println("From incom: ", moneyIncom, " rub  have to pay: ", progressiveTax, " rub")
	} else {
		progressiveTax = moneyIncom * 13 / 100
		fmt.Println("From incom: ", moneyIncom, " rub  have to pay: ", progressiveTax, " rub")
	}
}
