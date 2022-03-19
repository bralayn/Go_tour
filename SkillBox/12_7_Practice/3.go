//Задание 3. Уровни доступа
//
//Что нужно сделать
//Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.
//
//Рекомендация
//Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из него данные.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var text string
	someText := "May Peace prevail on Earth"
	//------------------------------------------------------------
	file, err := os.Create("checkFile.txt")
	if err != nil {
		fmt.Println("Can't create file", err)
		return
	}
	defer file.Close()

	file.WriteString(someText)
	fmt.Println("File", file.Name(), "was created")
	//-------------------------------------------------------------

	f, err := os.Open("checkFile.txt")
	if err != nil {
		panic(err)
	}

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	buf := make([]byte, fi.Size())
	if _, err := io.ReadFull(f, buf); err != nil {
		fmt.Println("Can't read file", err)
		return
	}
	fmt.Printf(" %s\n", buf)
	//----------------------------------------------------------------
	// Getting the current permissions of a file
	fmt.Printf("Current permissions - %v\n", fi.Mode())

	// Changing permissions
	err = file.Chmod(0444)
	if err != nil {
		panic(err)
	}
	fi, err = file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Changed NEW permissions - %v\n", fi.Mode())
	file.Close() // Close file checkFile.txt

	//--------------------------------------------------------------------
	fmt.Print("Input some text to ADD -> ")
	_, _ = fmt.Scan(&text)
	file.WriteString(text)

	//-------------------------------------------------------------------
	fn, err := os.Open("checkFile.txt")
	if err != nil {
		panic(err)
	}

	fin, err := fn.Stat()
	if err != nil {
		panic(err)
	}

	bufn := make([]byte, fin.Size())
	if _, err := io.ReadFull(fn, bufn); err != nil {
		fmt.Println("Can't read file", err)
		return
	}
	fmt.Printf("File after ADD text ->  %s\n", bufn)

}
