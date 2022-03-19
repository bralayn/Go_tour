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
	dtFormat := dt.Format("01-02-2006 15:04:05")
	fmt.Print("Input some text (if exit-end programm) -> ")
	i := 1

	for {
		_, _ = fmt.Scan(&text)
		if text == "exit" {
			fmt.Println("--- Exit programm ---")
			break
		}
		n := strconv.Itoa(i)
		file.WriteString(n)
		file.WriteString(fmt.Sprintf(" %s %s \n", dtFormat, text))
		i++
	}
	fmt.Println("Execution results in a file -> ", file.Name())
}
