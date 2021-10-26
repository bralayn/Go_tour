package main

import "fmt"

func main() {

	var (
		standartMenu  string = "Tea, sandwich with cheese and apple !!!"
		mondayMenu    string = "Pea soup, pasta with sausage +"
		tuesdayMenu   string = "Borsch, buckwheat porridge with chicken +"
		wednesdayMenu string = "Cabbage soup, grated potato casserole with fish +"
		thursdayMenu  string = "Fish soup, pearl barley porridge with cutlets +"
		fridayMenu    string = "Pickle and lasagna +"
		saturdayMenu  string = "Milk soup and rice porridge +"
		sundayMenu    string = "fried potatoes with meat and 300 grams of vodka +"
		day           int
	)
	fmt.Print("Input the day of the week from 1 to 7: ")
	fmt.Scan(&day)
	fmt.Println()

	if day == 1 {
		fmt.Println("--- Today monday ----")
		fmt.Println("----- Your menu for today -----")
		fmt.Println(mondayMenu, standartMenu)
	} else if day == 2 {
		fmt.Println("--- Today tuesday ----")
		fmt.Println("----- Your menu for today -----")
		fmt.Println(tuesdayMenu, standartMenu)
	} else if day == 3 {
		fmt.Println("--- Today wednesday ----")
		fmt.Println("----- Your menu for today -----")
		fmt.Println(wednesdayMenu, standartMenu)
	} else if day == 4 {
		fmt.Println("--- Today thursday ----")
		fmt.Println("----- Your menu for today -----")
		fmt.Println(thursdayMenu, standartMenu)
	} else if day == 5 {
		fmt.Println("--- Today friday ----")
		fmt.Println("----- Your menu for today -----")
		fmt.Println(fridayMenu, standartMenu)
	} else if day == 6 {
		fmt.Println("--- Today saturday ----")
		fmt.Println("----- Your menu for today -----")
		fmt.Println(saturdayMenu, standartMenu)
	} else {
		fmt.Println("--- Today sunday !!!! ----")
		fmt.Println("----- Your menu for today -----")
		fmt.Println(sundayMenu, standartMenu)
	}

}
