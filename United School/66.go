package main

import "fmt"

func four(s []int) {
	s = append(s, 4)
}

func main() {
	s := []int{1, 2, 3}
	four(s)
	fmt.Println(s)
}
