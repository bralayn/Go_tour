package main

import "fmt"

func main() {
	var sum, hoarder, n int

	/*
	 * Когда вы создали переменную sum она уже инициализирована
	 * нулевым значением этого типа, т.е. 0, это не ошибка
	 * но можно легко опустить строку:
	 * sum = 0
	 */

	/*
	 * А здесь вы пропустили переменную n, ее же нужно прочитать из консоли
	 * fmt.Scan(&n)
	 */

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&hoarder)
		if hoarder%8 == 0 && hoarder < 100 && hoarder > 9 {
			sum = sum + hoarder
		}
	}

	fmt.Print(sum)
}
