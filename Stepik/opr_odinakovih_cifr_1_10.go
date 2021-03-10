package main

import "fmt"

func main() {
	var (
		numX, numY, x1, x2, x3, x4,
		y1, y2, y3, y4 int
	)

	numX = 8
	numY = 8954
	//fmt.Scan(&numX, &numY)

	x1 = numX / 1000 % 10
	x2 = numX / 100 % 10
	x3 = numX / 10 % 10
	x4 = numX % 10

	y1 = numY / 1000 % 10
	y2 = numY / 100 % 10
	y3 = numY / 10 % 10
	y4 = numY % 10

	if x1 > 0 {

		switch x1 {
		case y1:
			fmt.Print(x1, " ")
		case y2:
			fmt.Print(x1, " ")
		case y3:
			fmt.Print(x1, " ")
		case y4:
			fmt.Print(x1, " ")
		}
	}

	if x1 == 0 && x2 == 0 {
		goto met1
	} else {
		switch x2 {
		case y1:
			fmt.Print(x2, " ")
		case y2:
			fmt.Print(x2, " ")
		case y3:
			fmt.Print(x2, " ")
		case y4:
			fmt.Print(x2, " ")
		}
	}
met1:
	if x1 == 0 && x2 == 0 && x3 == 0 {
		goto met2
	} else {

		switch x3 {
		case y1:
			fmt.Print(x3, " ")
		case y2:
			fmt.Print(x3, " ")
		case y3:
			fmt.Print(x3, " ")
		case y4:
			fmt.Print(x3, " ")
		}
	}
met2:

	switch x4 {
	case y1:
		fmt.Print(x4, " ")
	case y2:
		fmt.Print(x4, " ")
	case y3:
		fmt.Print(x4, " ")
	case y4:
		fmt.Print(x4, " ")
	}
}

//  Normalini sposob cherez  - strng -
//package main
//
//import "fmt"
//
//func main() {
//	var a, b string
//	fmt.Scan(&a, &b)
//
//	for i:=0 ; i < len(a) ; i++ {
//		for j:=0 ; j < len(b) ; j++ {
//			if a[i] == b[j] {
//				fmt.Print(string(a[i]), " ")
//				break     // цифры в числах не повторяются
//			}
//		}
//	}
//}

// --------   3 sposob -------------------
//package main
//import "fmt"
//
//func main()  {
//	var n, m, d int
//
//	fmt.Scan(&n, &m)
//	for i := 1; i <= n; i *= 10 {
//		d = n % (i * 10) / i
//		for j := 1; j <= m; j *= 10 {
//			if d == m % (j * 10) / j {
//				defer fmt.Print(d, " ")
//				break
//			}
//		}
//	}
//}

//	fmt.Println(x1, x2,x3, x4)
//	fmt.Println(y1, y2, y3, y4)
