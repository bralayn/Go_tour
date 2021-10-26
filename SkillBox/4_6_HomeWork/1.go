//Напишите программу, которая запрашивает у пользователя результаты ЕГЭ по трём экзаменам и проверяет,
//поступил ли он в институт или нет. Сумма проходных баллов в институт — 275

package main

import "fmt"

func main() {

	var rezultExam1, rezultExam2, rezultExam3 int

	needBalls := 275

	fmt.Println("-- points unified state examination --")
	fmt.Print("Enter the result of the first exam: ")
	fmt.Scan(&rezultExam1)
	fmt.Println()

	fmt.Print("Enter the result of the second exam: ")
	fmt.Scan(&rezultExam2)
	fmt.Println()

	fmt.Print("Enter the result of the third exam: ")
	fmt.Scan(&rezultExam3)
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
