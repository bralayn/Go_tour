package main

import "fmt"

func main() {
	var f float64
	fmt.Print("Input temperaturu v Forengeitah =>")
	fmt.Scanf("%f", &f)
	c := (f - 32.0) * 5.0 / 9.0
	fmt.Println("Temperatura v Celsiyah = ", c)
}
