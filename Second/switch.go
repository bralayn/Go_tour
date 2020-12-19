package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on - ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Os X.")
	case "Linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)

	}
}
