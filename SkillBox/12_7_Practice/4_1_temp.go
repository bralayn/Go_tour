//Задание 4_1. Работа с файлами
//---Перепишите задачи 1 используя пакет ioutil
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var text, n, ss string
	var buf bytes.Buffer

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
		n = strconv.Itoa(i)
		ss = fmt.Sprintf("%s %s %s \n", n, dtFormat, text)

		r := strings.NewReader(ss)

		b, err := ioutil.ReadAll(r)
		if err != nil {
			panic(err)
		}
		buf.WriteString(fmt.Sprintf("%s", b))
		i++
	}
	fileName := "logFile.txt"
	if err := ioutil.WriteFile(fileName, buf.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Execution results in a file -> ", file.Name())
	fmt.Println(string(resBytes))
}
