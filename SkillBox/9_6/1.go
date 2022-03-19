//Задание 1. Переполнение
//
//Что нужно сделать
//В данном модуле мы рассмотрели примеры по целочисленным типам, их размерам в памяти и то, что происходит
//при её переполнении. Напишите программу, которая в цикле с использованием встроенных констант
//(на предельные значения целых чисел, в пакете math) будет подсчитывать, сколько приходится переполнений чисел
//типа uint8, uint16 в диапазоне от 0 до uint32.
//
//Советы и рекомендации
//Для нахождения количества переполнений используйте цикл. Воспользуйтесь константами максимальных значений из пакета math.
package main

import (
	"fmt"
	"unsafe"
)

var i, countUint8, countUint16 int

const (
	toBe      bool   = false
	numUint8  uint8  = 1<<8 - 1
	numUint16 uint16 = 1<<16 - 1
	maxUint32 uint32 = 1<<32 - 1
)

func main() {
	fmt.Printf("Type: %T Size: %d byte Value: %v\n", toBe, unsafe.Sizeof(toBe), toBe)
	fmt.Printf("Type: %T Size: %d byte Value: %v\n", numUint8, unsafe.Sizeof(numUint8), numUint8)
	fmt.Printf("Type: %T Size: %d byte Value: %v\n", numUint16, unsafe.Sizeof(numUint16), numUint16)
	fmt.Printf("Type: %T Size: %d byte Value: %v\n", maxUint32, unsafe.Sizeof(maxUint32), maxUint32)
	fmt.Println("==================================================================")

	for i = 0; uint32(i) <= maxUint32; i++ {
		if uint32(i)%uint32(numUint8) == 0 {
			countUint8 += 1
			fmt.Println("Count Uint8 = ", countUint8)
		}
		if uint32(i)%uint32(numUint16) == 0 {
			countUint16 += 1
			fmt.Println("Count Uint16 = ", countUint16)
		}
	}
	fmt.Println("Total Uint 8=", countUint8)
	fmt.Println("Total Uint16=", countUint16)
}
