package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz i= ", i)
		case i%3 == 0:
			fmt.Println("Fizz i= ", i)
		case i%5 == 0:
			fmt.Println("Buzz i= ", i)

		}
	}
}
