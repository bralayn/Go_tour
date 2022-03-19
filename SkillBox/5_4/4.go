//Задание 4. Сумма без сдачи
//
//Что нужно сделать
//Программное обеспечение банкоматов постоянно решает задачу, как имеющимися купюрами сформировать сумму,
//введённую пользователем. Попробуйте решить похожую задачу и определить, сможет ли пользователь заплатить
//за товар без сдачи или нет. Для этого он будет вводить стоимость товара и номиналы трёх монет.
//
//Рекомендация
//Проверьте все сценарии, когда сумма может быть сформирована одной монетой, двумя или всеми тремя.
package main

import "fmt"

func main() {

	var coin1, coin2, coin3, costOfGoods, sumCoins int

	fmt.Print("Input cost of goods: ")
	fmt.Scan(&costOfGoods)
	fmt.Println()

	fmt.Print("Input coin1: ")
	fmt.Scan(&coin1)
	fmt.Println()
	if costOfGoods == coin1 {
		fmt.Println("1 coin=", coin1, " is enough to pay for the goods ")
		goto endPrg
	}

	fmt.Print("Input coin2: ")
	fmt.Scan(&coin2)
	fmt.Println()
	if costOfGoods == (coin1 + coin2) {
		fmt.Println("2 coins=", coin1, "+", coin2, " is enough to pay for the goods ")
		goto endPrg
	}

	fmt.Print("Input coin3: ")
	fmt.Scan(&coin3)
	fmt.Println()
	sumCoins = coin1 + coin2 + coin3
	if costOfGoods == sumCoins {
		fmt.Println("3 coins=", coin1, "+", coin2, "+", coin3, " is enough to pay for the goods ")
	}

	if costOfGoods < sumCoins || costOfGoods < (coin1+coin2) || costOfGoods < coin1 {
		fmt.Println("Need to give change ")
	}
	if costOfGoods > sumCoins {
		fmt.Println("Not enough money!!! Goodbye my friend.")
	}

endPrg:
}
