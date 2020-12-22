package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// Проверяем, какие числа деляться на 3, 6, 9 , те числа и выводим на экран
		if i%3 == 0 && i%6 == 0 && i%9 == 0 {
			fmt.Println(i)
		}
	}
}
