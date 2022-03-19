//Задание 2. Вывод чисел, содержащихся в строке
//
//Что нужно сделать
//Напишите программу, которая выведет все части строки
//
//a10 10 20b 20 30c30 30 dd,
//
//которые можно привести к числу в десятичном формате.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "a10 10 20b 20 30c30 30 dd"
	fmt.Println("Исходная строка:")
	fmt.Println(s)
	fmt.Println("В строке содержатся числа в десятичном формате:")
	s = s + " "
	for len(s) > 0 {
		spaceIndex := strings.Index(s, " ")
		firstNumber := s[:spaceIndex]

		if firstNumber, err := strconv.Atoi(firstNumber); err == nil {
			fmt.Print(firstNumber, " ")
		}

		s = s[spaceIndex+1:]
	}
}
