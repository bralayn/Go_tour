//Что нужно сделать
//Напишите программу, которая запрашивает день недели, число гостей и сумму чека и рассчитывает сумму к оплате.
//В ресторане действуют следующие правила:
//по понедельникам должна применяться скидка 10% на всё меню, потому что понедельник — день тяжёлый;
//по пятницам, если сумма чека превышает 10 000 рублей, включается дополнительная скидка в размере 5%;
//если число гостей в одной компании превышает пять человек, автоматически включается надбавка на обслуживание 10%.

package main

import "fmt"

func main() {

	var (
		numberDay, numberOfguests, checkAmount        int
		mondayDiscount, bigsumDiscount, peoplDiscount int
	)

	fmt.Print("Input the Day of the Week from 1 to 7: ")
	fmt.Scan(&numberDay)
	fmt.Println()

	fmt.Print("Input number of guests: ")
	fmt.Scan(&numberOfguests)
	fmt.Println()

	fmt.Print("Input check amount: ")
	fmt.Scan(&checkAmount)
	fmt.Println()

	mondayDiscount = checkAmount * 10 / 100
	bigsumDiscount = checkAmount * 5 / 100
	peoplDiscount = checkAmount * 10 / 100

	if numberDay == 1 {
		if numberOfguests > 5 {
			fmt.Println("Discount on Mondays is 10% : ", mondayDiscount)
			fmt.Println("Service surcharge is + 10% :", peoplDiscount)
			fmt.Println("Amount to pay:", checkAmount-mondayDiscount+peoplDiscount)
		} else {
			fmt.Println("Discount on Mondays is 10% : ", mondayDiscount)
			fmt.Println("Amount to pay:", checkAmount-mondayDiscount)
		}
	}

	if numberDay == 5 {
		if checkAmount > 10000 {
			if numberOfguests > 5 {
				fmt.Println("Discount on Friday is 5% : ", bigsumDiscount)
				fmt.Println("Service surcharge is + 10% :", peoplDiscount)
				fmt.Println("Amount to pay:", checkAmount-bigsumDiscount+peoplDiscount)
			} else {
				fmt.Println("Discount on Friday is 5% : ", bigsumDiscount)
				fmt.Println("Amount to pay:", checkAmount-bigsumDiscount)
			}
		} else {
			fmt.Println("Amount to pay:", checkAmount)
		}
	} else if numberOfguests > 5 {
		fmt.Println("Service surcharge is + 10% :", peoplDiscount)
		fmt.Println("Amount to pay:", checkAmount+peoplDiscount)
	}
}
