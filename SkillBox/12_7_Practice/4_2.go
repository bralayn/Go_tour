//Задание 4_2. Интерфейс io.Reader
//
//Что нужно сделать
//Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике,
//используя ioutil.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("logFile.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	resBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resBytes))
}
