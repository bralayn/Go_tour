//Задание:
//
//В начале программы объявите строку и присвойте ей:
//Golang is programming language
//
//Напишите код, который посчитает количество "a" в строке.
//Можно использовать ReplaceAll (с заменой на пустую строку), чтобы сравнить длину исходной строки и длину строки после замены.
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Golang is programming language"
	s1 := strings.ReplaceAll(s, "a", "")
	fmt.Println(s1)
	fmt.Println(len(s))
	fmt.Println(len(s1))
	fmt.Println("Кол-во 'a' в строке s=", len(s)-len(s1))
}
