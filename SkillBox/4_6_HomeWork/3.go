//Что нужно сделать
//Банкомат выдаёт только купюры номиналом 100 рублей, а максимальная сумма снятия — 100 000 рублей.
//Напишите программу, которая проверяет допустимость введённой пользователем суммы средств.
//Если сумма для снятия доступна, сообщите «Операция успешно выполнена», в ином случае укажите причину,
//по которой невозможно выполнить операцию.

package main

import "fmt"

func main() {
	var needCash, nominalCash int

	nominalCash = 100

	fmt.Println("--- ATM ---")
	fmt.Print("Enter the amount to be withdrawn from the account: ")
	fmt.Scan(&needCash)
	fmt.Println()

	if needCash <= 100000 {

		if needCash%nominalCash == 0 {
			fmt.Println("Operation completed successfully")
			fmt.Println("You have withdrawn", needCash, " rubles ")
		} else {
			fmt.Println("Enter the correct amount multiple of 100 ")
		}
	} else {
		fmt.Println("Withdrawal amount must not exceed 100.000 ")
	}

}
