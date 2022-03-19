//Задание 1. Определение количества слов, начинающихся с большой буквы
//
//
//Что нужно сделать
//Напишите программу, которая выведет количество слов, начинающихся с большой буквы в строке: Go is an Open source
//programming Language that makes it Easy to build simple, reliable, and efficient Software.
package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
	s = s + " "
	count := 0
	for len(s) > 0 {
		spaceIndex := strings.Index(s, " ")
		firstWord := s[:spaceIndex]
		if firstWord == strings.Title(firstWord) {
			count++
		}
		s = s[spaceIndex+1:]

	}
	fmt.Println("Определение количества слов, начинающихся с большой буквы в строке:")
	fmt.Println("Go is an Open source programming Language that makes it Easy")
	fmt.Println(" to build simple, reliable, and efficient Software.")
	fmt.Println("Строка содержит", count, "слов с большой буквы")
}
