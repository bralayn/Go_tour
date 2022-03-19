//Задание 2. Шахматная доска
//
//Что нужно сделать
//Одним из видов компьютерного искусства является псевдографика, когда из символов создаются картины.
//Попробуйте вывести два изображения. Запросите у пользователя размер шахматной доски в клетках и выведите
//шахматную доску на экран. Белые поля выведите звёздочкой, а чёрные — пробелом.
//
//Советы и рекомендации
//Можно использовать fmt.Print и fmt.Println.
//
//Пример работы программы:
//
//Шахматная доска.
//Введите ширину:
//5
//Введите высоту:
//4
// * *
//* * *
// * *
//* * *
package main

import "fmt"

func main() {
	var i, j, heightBoard, widthBoard int

	fmt.Print("Input width =")
	fmt.Scan(&widthBoard)
	fmt.Println()

	fmt.Print("Input height =")
	fmt.Scan(&heightBoard)
	fmt.Println()

	i = 1
	j = 1

	for j <= heightBoard {

		for i <= widthBoard {
			if (i+j)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
			i++
		}
		fmt.Println()
		j++
		i = 1
	}
}
