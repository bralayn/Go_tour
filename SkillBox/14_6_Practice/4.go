//Задание 4. Область видимости переменных
//
//Что нужно сделать
//Напишите программу, в которой будет три функции, попарно использующие разные глобальные переменные.
//Функции должны прибавлять к поданному на вход числу глобальную переменную и возвращать результат.
//Затем вызовите по очереди три функции, передавая результат из одной в другую.
package main

import "fmt"

var z1, z2, z3 = 55, 66, 77

func f1(x int) (x1 int) {
	x1 = x + z1
	return
}

func f2(y int) (y1 int) {
	y1 = y + z2
	return
}

func f3(n int) (n1 int) {
	n1 = n + z3
	return
}
func main() {
	fmt.Printf("Global value z1=%d z2=%d z3=%d\n", z1, z2, z3)

	fmt.Println("f1=", f1(10))
	fmt.Println("f2=", f2(15))
	fmt.Println("f3=", f3(20))

	sumF1F2 := f1(10) + f2(15)
	fmt.Println("Summa f1 + f2 =", sumF1F2)

	rez := f1(f2(10))
	fmt.Println("Rezult f1(f2) =", rez)

	anotherRez := f3(f2(f1(30)))
	fmt.Println("Another example f3(f2(f1)) =", anotherRez)

	fmt.Println("Example f2- f3(f2(f1))=", f2(50)-anotherRez)
}
