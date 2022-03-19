package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	answer := "да"

	for answer == "да" {
		a := rand.Intn(9) + 1
		b := rand.Intn(9) + 1
		//	fmt.Println(a)
		//	fmt.Println(b)
		result := 0
		for a*b != result {
			fmt.Print(a, "*", b, "=")
			fmt.Scan(&result)
			if a*b != result {
				fmt.Println("Ответ не правильный, попробуйте еще раз")
			}
		}
		fmt.Println("Ответ правильный")
		fmt.Println("Еще пример (да/нет)? ")
		fmt.Scan(&answer)
		for answer != "да" && answer != "нет" {
			fmt.Println("Еще пример (да/нет)? ")
			fmt.Scan(&answer)
		}
	}
}
