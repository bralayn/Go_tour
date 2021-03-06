package main

import "fmt"

func main() {
	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:5]

	sr := []float64{1, 2, 3, 4, 5, 6}
	y := sr[0:5]

	fmt.Println(x)
	fmt.Println(y)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice11 := []int{1, 2, 3}
	slice22 := make([]int, 2)
	copy(slice22, slice11)
	fmt.Println(slice11, slice22)
}
