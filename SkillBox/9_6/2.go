//Задание 2. Минимальный тип данных
//
//Что нужно сделать
//Достаточно часто при передаче по Сети или сохранении больших объёмов данных приходится выбирать тип с минимальным
//размером памяти, чтобы экономить трафик или место на диске. Напишите программу, в которую пользователь вводит
//два числа (int16), а программа выводит, в какой минимальный тип данных можно сохранить результат умножения этих чисел.
//
//Советы и рекомендации
//Обратите внимание, что положительный результат можно сохранить в меньшем типе за счёт использования uint8, uint16.
//Чтобы не возникло проблем с переполнением в процессе умножения, числа считывайте в int16,
//а перед умножением переведите их в int32.
//
//Проверить на примерах:
//
//1 1 результат uint8
//
//1 −1 результат int8
//
//640 100 результат uint16
//
//−640 100 результат int32
//
//3000 3000 результат uint32
//
//−3000 3000 результат int32

/*
--    for myself ------

int8     -128 / 127
uint8       0 / 255
int16  -32768 / 32767
uint16      0 / 65535
int32  -2147483648 / 2147483647
uint32      0 / 4294967295

*/

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var (
		x1, x2   int16
		rezult   int32
		maxInt16 int16 = 1<<15 - 1
		maxInt32 int32 = 1<<31 - 1
	)
	fmt.Printf("Type: %T Size: %d byte Value: %v\n", maxInt16, unsafe.Sizeof(maxInt16), maxInt16)
	fmt.Printf("Type: %T Size: %d byte Value: %v\n", maxInt32, unsafe.Sizeof(maxInt32), maxInt32)

	fmt.Print("Input first number=")
	fmt.Scan(&x1)
	fmt.Println()
	fmt.Print("Input second number=")
	fmt.Scan(&x2)
	fmt.Println()

	rezult = int32(x1) * int32(x2)

	if rezult >= -128 && rezult < 0 {
		fmt.Println("Rezult Int8 =", rezult)
		goto endPrg
	}
	if rezult >= 0 && rezult <= 255 {
		fmt.Println("Rezult Uint8 =", rezult)
		goto endPrg
	}
	if rezult >= -32768 && rezult < 0 {
		fmt.Println("Rezult Int16 =", rezult)
		goto endPrg
	}
	if rezult >= 0 && rezult <= 65535 {
		fmt.Println("Rezult Uint16 =", rezult)
		goto endPrg
	}
	if rezult >= -2147483648 && rezult < 0 {
		fmt.Println("Rezult Int32 =", rezult)
		goto endPrg
	}
	if uint32(rezult) >= 0 && uint32(rezult) <= 4294967295 {
		fmt.Println("Rezult Uint32 =", rezult)
		goto endPrg
	}

endPrg:
}
