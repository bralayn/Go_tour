package main

import "fmt"

func main() {
	var (
		firstEmployee, secondEmployee, thirdEmployee        int
		minSalary, middlSalary, maxSalary, differenceSalary int
	)
	fmt.Print("Input the first employee salary: ")
	fmt.Scan(&firstEmployee)
	fmt.Println()

	fmt.Print("Input the second employee salary: ")
	fmt.Scan(&secondEmployee)
	fmt.Println()

	fmt.Print("Input the third employee salary: ")
	fmt.Scan(&thirdEmployee)
	fmt.Println()

	//  Используем метод Пузырьковая сортировка

	if firstEmployee > secondEmployee {
		maxSalary = firstEmployee
		minSalary = secondEmployee
	} else {
		maxSalary = secondEmployee
		minSalary = firstEmployee
	}

	if maxSalary > thirdEmployee {
		fmt.Println("MaxSalary :", maxSalary)
	} else {
		maxSalary = thirdEmployee
		fmt.Println("MaxSalary :", maxSalary)
	}

	if minSalary < thirdEmployee {
		fmt.Println("MinSalary :", minSalary)
	} else {
		minSalary = thirdEmployee
		fmt.Println("MinSalary :", minSalary)
	}
	differenceSalary = maxSalary - minSalary
	middlSalary = (firstEmployee + secondEmployee + thirdEmployee) / 3

	fmt.Println("MiddlSalary = ", middlSalary)
	fmt.Println("Difference salary =", differenceSalary)

}
