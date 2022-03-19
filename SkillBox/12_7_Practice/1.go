//Задание 1. Работа с файлами
//
//Что нужно сделать
//Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату
//и сообщение в формате:
//
//2020-02-10 15:00:00 продам гараж.
//
//При вводе слова exit программа завершает работу.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var text string

	file, err := os.Create("logFile.txt")
	if err != nil {
		fmt.Println("Can't create file", err)
		return
	}
	defer file.Close()

	dt := time.Now()
	fmt.Print("Input some text -> ")

	for i := 1; i < 10; i++ {
		_, _ = fmt.Scan(&text)
		//	fmt.Scan(&text)
		if text == "exit" {
			fmt.Println("Exit programm")
			break
		}
		file.WriteString(strconv.Itoa(i))
		//file.WriteString(". ")
		//file.WriteString(text)
		//file.WriteString(t.Format("01-02-2003 15:04:05"))
		file.WriteString(fmt.Sprintf(". %s %s \n", dt.Format("01-02-2003 15:04:05"), text))
		//	file.WriteString("\n")
		//i++

		fmt.Println(dt.Format("01-02-2003 15:04:05"))
	}
	fmt.Println(file)
	fmt.Println(file.Name())
}
