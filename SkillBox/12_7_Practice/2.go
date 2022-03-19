//Задание 2. Интерфейс io.Reader
//
//Что нужно сделать
//Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике,
//без использования ioutil. Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.
//
//Рекомендация
//Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("logFile.txt")
	if err != nil {
		fmt.Println("File open error ")
		//panic(err)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Size() != 0 {
		buf := make([]byte, fi.Size())
		if _, err := io.ReadFull(f, buf); err != nil {
			fmt.Println("Can't read file", err)
			return
		}
		fmt.Printf(" %s\n", buf)
		fmt.Println(string(buf)) // Another metod

	} else {
		fmt.Println("File is empty ")
	}

	fmt.Printf("Size: %d\n", fi.Size())
}
