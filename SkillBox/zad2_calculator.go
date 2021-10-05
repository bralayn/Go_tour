package main

import "fmt"

func main() {

	priceProduct := 5800
	priceDelivery := 430
	discount := 600
	priceEnd := priceProduct + priceDelivery - discount

	fmt.Println("Стоимость товара:", priceProduct)
	fmt.Println("Стоимость доставки:", priceDelivery)
	fmt.Println("Размер скидки:", discount)
	fmt.Println("---------------")
	fmt.Println("Итого:", priceEnd)
}
