package main

import "fmt"

func main() {

	var firstAstronaut, secondAstronaut, thirdAstronaut, maxIQ int

	fmt.Print("Input IQ first astronaut: ")
	fmt.Scan(&firstAstronaut)
	fmt.Println()

	fmt.Print("Input IQ second astronaut: ")
	fmt.Scan(&secondAstronaut)
	fmt.Println()

	fmt.Print("Input IQ third astronaut: ")
	fmt.Scan(&thirdAstronaut)
	fmt.Println()

	if firstAstronaut > secondAstronaut {
		maxIQ = firstAstronaut
		fmt.Println("Max IQ first")
	} else {
		maxIQ = secondAstronaut
		fmt.Println("Max IQ second")
	}
	if maxIQ < thirdAstronaut {
		maxIQ = thirdAstronaut
		fmt.Println("Max IQ third")
	}
	fmt.Println("The ommander will with IQ= ", maxIQ)

}
