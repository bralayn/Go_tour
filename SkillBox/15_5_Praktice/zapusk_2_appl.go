package main

import "fmt"

var a int

func changeA(b int) {
	a = b
}

func main() {
	for i := 0; i < 10; i++ {
		changeA(i)
		fmt.Println(a)
	}
}
