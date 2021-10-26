package main

import "fmt"

func main() {
	var (
		totalHeight, day, matureBambuk int
		seedlingBambuk                 int
		growBambuk                     int
		eatBambuk                      int
	)

	fmt.Print("Введите высоту саженца бамбука в см: ")
	fmt.Scan(&seedlingBambuk)
	fmt.Println()

	fmt.Print("Введите рост бамбука за день в см:")
	fmt.Scan(&growBambuk)
	fmt.Println()

	fmt.Print("Введите сколько съедают гусеницы за ночь в см:")
	fmt.Scan(&eatBambuk)
	fmt.Println()

	fmt.Print("Введите желаемую высоту созревания бамбука в см:")
	fmt.Scan(&matureBambuk)
	fmt.Println()

	totalHeight = seedlingBambuk

l1:
	if totalHeight <= matureBambuk {
		totalHeight += growBambuk
		day += 1
		fmt.Println(totalHeight, day)
		if totalHeight < matureBambuk {
			totalHeight -= eatBambuk
			fmt.Println(totalHeight, day)
		}
	}
	if totalHeight < matureBambuk {
		goto l1
	}
	fmt.Println("----------------- Результат работы программы -------------------")
	fmt.Println("Срубить и продать бамбук можно будет через", day, "полных дней.")
}
