//Задание 5. Определение максимальных процентов
//
//Что нужно сделать
//Напишите программу, которая позволит пользователю ввести три процентные ставки и выведет
//в консоль две наиболее выгодные, а если все процентные ставки равны — сообщит, что ставки равнозначны.
//
//Рекомендация
//Задачу можно решить несколькими способами — например, от противного.
package main

import "fmt"

func main() {

	var interestRate1, interestRate2, interestRate3, maxRate, middlRate float64

	fmt.Println("--- Enter three bank interest rates ---")

	fmt.Print("Interest rate 1=")
	fmt.Scan(&interestRate1)
	fmt.Println()

	fmt.Print("Interest rate 2=")
	fmt.Scan(&interestRate2)
	fmt.Println()

	fmt.Print("Interest rate 3=")
	fmt.Scan(&interestRate3)
	fmt.Println()

	if interestRate1 == interestRate2 && interestRate2 == interestRate3 {
		fmt.Println("--- All interest rates are equal !!! ---")
		goto endPrg
	}

	if interestRate1 > interestRate2 && interestRate1 > interestRate3 {
		maxRate = interestRate1
		if interestRate2 > interestRate3 {
			middlRate = interestRate2
		} else {
			middlRate = interestRate3
		}
	} else if interestRate2 > interestRate1 && interestRate2 > interestRate3 {
		maxRate = interestRate2
		if interestRate1 > interestRate3 {
			middlRate = interestRate1
		} else {
			middlRate = interestRate3
		}
	} else {
		maxRate = interestRate3
		if interestRate1 > interestRate2 {
			middlRate = interestRate1
		} else {
			middlRate = interestRate2
		}
	}
	fmt.Println("--------- Results ----------")
	fmt.Println("Max rate =", maxRate)
	fmt.Println("Middle rate =", middlRate)

endPrg:
}
