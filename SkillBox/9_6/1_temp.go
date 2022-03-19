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
	numUint8  = 255
	numUint16 = 65535
	maxUint32 = 4294967295
)

func main() {
	fmt.Printf("Type: %T Size: %d byte Value: %v\n", numUint8, unsafe.Sizeof(numUint8), numUint8)
	fmt.Printf("Type: %T Size: %d byte Value: %v\n", numUint16, unsafe.Sizeof(numUint16), numUint16)
	fmt.Printf("Type: %T Size: %d byte Value: %v\n", maxUint32, unsafe.Sizeof(maxUint32), maxUint32)
	fmt.Println("==================================================================")

	for i <= maxUint32 {

		if i%numUint8 == 0 {
			countUint8 = countUint8 + 1
			fmt.Println("Uint8= ", countUint8)
			//continue
		}

		if i%numUint16 == 0 {
			countUint16 = countUint16 + 1
			fmt.Println("Uin16= ", countUint16)
			//continue
		}
		i++
	}
	fmt.Println("Uint 8", countUint8)
	fmt.Println("Uint16", countUint16)
}
