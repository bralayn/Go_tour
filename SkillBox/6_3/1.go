//Задача 1. Проверка пароля
//
//Напишите программу с бесконечным циклом, которая спрашивает у пользователя пароль,
//и совершает выход из цикла с помощью инструкции break.
package main

import "fmt"

func main() {
	var psw int

	for true {
		fmt.Print("Input password= ")
		fmt.Scan(&psw)
		fmt.Println()
		if psw == 666 {
			break
		}

	}
	fmt.Println("--- End program ---")
}
