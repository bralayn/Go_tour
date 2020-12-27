package main

import "fmt"

//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}   // 1-st variant declarated

func main() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128, 256} //2-d variant declarated
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
