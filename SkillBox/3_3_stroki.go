package main

import "fmt"

func main() {
	var (
		name, password string
		age            int
	)
	fmt.Print("Введите имя пользователя --> ")
	fmt.Scan(&name)
	fmt.Println()

	fmt.Print("Введите пароль --> ")
	fmt.Scan(&password)
	fmt.Println()

	fmt.Print("Введите Ваш возраст --> ")
	fmt.Scan(&age)
	fmt.Println()

	fmt.Println("Поздравляю", name, "Теперь вы зарегестрированы")
	fmt.Println("Ваш пароль:", password)
	fmt.Println("Ваш возраст:", age)
}
