//Напишите программу, которая запрашивает у пользователя результаты ЕГЭ по трём экзаменам и проверяет,
//поступил ли он в институт или нет. Сумма проходных баллов в институт — 275

// ----  Сделайте проверку на отрицательные значения и на лимит в 100 баллов. Жду эту правку:)

package main

import "fmt"

func main() {

	var rezultExam1, rezultExam2, rezultExam3 int

	needBalls := 275

	fmt.Println("-- points unified state examination --")
Exam1:
	fmt.Print("Enter the result of the first exam: ")
	fmt.Scan(&rezultExam1)
	if rezultExam1 < 0 {
		fmt.Println("!!! Enter the correct result > 0 ")
		goto Exam1
	} else if rezultExam1 > 100 {
		fmt.Println("The result cannot be more than 100, repeat please  ")
		goto Exam1
	}
	fmt.Println()

Exam2:
	fmt.Print("Enter the result of the second exam: ")
	fmt.Scan(&rezultExam2)
	if rezultExam2 < 0 {
		fmt.Println("!!! Enter the correct result > 0 ")
		goto Exam2
	} else if rezultExam2 > 100 {
		fmt.Println("The result cannot be more than 100, repeat please  ")
		goto Exam2
	}
	fmt.Println()

Exam3:
	fmt.Print("Enter the result of the third exam: ")
	fmt.Scan(&rezultExam3)
	if rezultExam3 < 0 {
		fmt.Println("!!! Enter the correct result > 0 ")
		goto Exam3
	} else if rezultExam3 > 100 {
		fmt.Println("The result cannot be more than 100, repeat please  ")
		goto Exam3
	}
	fmt.Println()

	allBalls := rezultExam1 + rezultExam2 + rezultExam3
	if allBalls < 275 {
		fmt.Println("The amount of passing points: ", needBalls)
		fmt.Println("Number of points scored: ", allBalls)
		fmt.Println("--- You failed the exam ---")
	} else {
		fmt.Println("The amount of passing points: ", needBalls)
		fmt.Println("Number of points scored: ", allBalls)
		fmt.Println("--- Congratulations. You entered!!! --- ")
	}

}
