//Задание 2. Функция, возвращающая несколько значений
//
//Что нужно сделать
//Напишите программу, которая с помощью функции генерирует три случайные точки в двумерном пространстве (две координаты)
//а затем с помощью другой функции преобразует эти координаты по формулам: x = 2 × x + 10, y = −3 × y − 5.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x1, y1 := random2Number(10, 20)
	fmt.Println("x1=", x1, "y1=", y1)
	changeRezult(x1, y1)

	x2, y2 := random2Number(20, 30)
	fmt.Println("x2=", x2, "y2=", y2)
	changeRezult(x2, y2)

	x3, y3 := random2Number(30, 40)
	fmt.Println("x3=", x3, "y3=", y3)
	changeRezult(x3, y3)

}

func random2Number(n1, n2 int) (int, int) {
	return rand.Intn(n1), rand.Intn(n2)
}

func changeRezult(x, y int) {
	x = 2*x + 10
	y = -3*y - 5
	fmt.Println("New koordinati", x, y)
}
