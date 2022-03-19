//Задание 3. Именованные возвращаемые значения
//
//Что нужно сделать
//Напишите программу, которая на вход получает число, затем с помощью двух функций преобразует его. Первая умножает,
//а вторая прибавляет число, используя именованные возвращаемые значения.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	num := getNumber(100)
	fmt.Println("First Number=", num)

	fmt.Println("Num * random number for multiply =", getMultiply(num))
	fmt.Println("Num + random number for addition =", getPlus(num))

}

func getNumber(n int) int {
	return rand.Intn(n)
}

func getMultiply(num int) (multNum int) {
	randomMultNum := getNumber(20)
	fmt.Println("Random number for multiply =", randomMultNum)
	multNum = num * randomMultNum
	return

}

func getPlus(num int) (sum int) {
	randomNum := getNumber(20)
	fmt.Println("Random number for addition =", randomNum)
	sum = num + randomNum
	return
}
