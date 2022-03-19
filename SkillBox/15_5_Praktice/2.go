package main

import "fmt"

func main() {
	var num [10]int

	num = enterArray()
	//for i := 9; i >= 0; i-- {
	//	fmt.Println(num[i])
	//}

	for i := 0; i < 5; i++ {
		num[i], num[9-i] = num[9-i], num[i]
	}
	fmt.Println(num)
}

func enterArray() (num [10]int) {
	fmt.Println("--- Input array numbers: ")
	for i := range num {
		fmt.Printf("Input number [%v] ", i+1)
		fmt.Scan(&num[i])
	}
	return num
}
