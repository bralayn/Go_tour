package main

import "fmt"

func main() {
	arr := [2]int{2, 6}
	fmt.Println(arr)
	arr[0], arr[1] = arr[1], arr[0]
	fmt.Println(arr)
}
