//Задание 6. Решение квадратного уравнения
//
//Что нужно сделать
//Что компьютеры делают быстрее людей? Выполняют вычисления!
//
//Напишите программу, решающую квадратные уравнения. Пользователь вводит коэффициенты a, b и c квадратного уравнения.
//Программа должна вывести корни уравнения. Для решения уравнения необходимо:
//
//Вычислить дискриминант по формуле: D = (b × b − 4 × a × c)
//
//Если D > 0, будет два различных корня, которые находятся по формуле.
//Если D = 0, будет один корень, который находится по формуле.
//Если D < 0, то вывести, что корней нет.
//
//Советы и рекомендации
//Используйте переменные типа float64.
//
//Для возведения b в квадрат воспользуйтесь функцией math.Pow.
//
//Для вычисления квадратного корня воспользуйтесь функцией math.Sqrt.
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64

	fmt.Print("Input a: ")
	fmt.Scan(&a)
	fmt.Print()
	fmt.Println()

	fmt.Print("Input b: ")
	fmt.Scan(&b)
	fmt.Print()
	fmt.Println()

	fmt.Print("Input c: ")
	fmt.Scan(&c)
	fmt.Print()
	fmt.Println()

	// calculate the discriminant and roots
	d = math.Pow(b, 2) - 4*a*c
	x1 := (-b + math.Sqrt(d)) / 2 * a
	x2 := (-b - math.Sqrt(d)) / 2 * a
	x := -b / 2 * a

	if d < 0 {
		fmt.Println(" No roots !!!")
	}

	if d > 0 {
		fmt.Println("--- equation solution if d > 0 ---")
		fmt.Println("x1=", x1, "x2=", x2)
	} else if d == 0 {
		fmt.Println("--- equation solution if d = 0 ---")
		fmt.Println("x=", x)
	}
}
