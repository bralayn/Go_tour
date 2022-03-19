package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		a := rand.Intn(9) + 1
		b := rand.Intn(9) + 1
		for {
			fmt.Print(a, "*", b, "=")
			var result int
			fmt.Scan(&result)
			if a*b != result {
				fmt.Println("Ответ не правильный, попробуйте еще раз")
			} else {
				break
			}
		}
		fmt.Println("Ответ правильный")
		fmt.Println("Еще пример (да/нет)? ")
		var answer string
		fmt.Scan(&answer)
		for answer != "да" && answer != "нет" {
			fmt.Println("Еще пример (да/нет)? ")
			fmt.Scan(&answer)
		}
		if answer == "нет" {
			break
		}
	}
}
