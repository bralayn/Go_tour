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
		if maxIQ > thirdAstronaut {
			fmt.Println("Max IQ first", maxIQ)
		}
	}

	if secondAstronaut > firstAstronaut {
		maxIQ = secondAstronaut
		if maxIQ > thirdAstronaut {
			fmt.Println("Max IQ second")
		}
	}

	if thirdAstronaut > secondAstronaut {
		maxIQ = thirdAstronaut
		if maxIQ > firstAstronaut {
			fmt.Println("Max IQ third")
		}
	}

	fmt.Println("The ommander will with IQ= ", maxIQ)

}
