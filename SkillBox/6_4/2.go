//Задача 2. Часовой
//
//Часовому нужно проходить вдоль КПП 20 метров. По 10 метров в каждую сторону от КПП. Своеобразная ось X.
//Напишите программу с безусловным циклом, которая будет считать шаги часового от -10 до 10 и выводить сообщение
//каждый раз при переходе через КПП, который на графике находите в 0 координате.
package main

import "fmt"

func main() {

	for {

		for i := -10; i <= 10; i++ {
			if i == 0 {
				fmt.Println("-- Prohodim KPP --")
			}
		}
		for j := 10; j >= -10; j-- {
			if j == 0 {
				fmt.Println("-- Prohodim KPP again")
			}
		}
	}

}
