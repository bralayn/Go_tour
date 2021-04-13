package main

import "fmt"

func main() {

	//workArray:= [10]uint8{}
	//idx:= [6]int{}
	//
	//var a uint8
	//var j int=0
	//for fmt.Scan(&a);j<16;fmt.Scan(&a){
	//	if j >=10{
	//		idx[j-10]= int(a)
	//	}else{
	//		workArray[j]=a
	//	}
	//	j++
	//}
	//fmt.Println(workArray)
	//fmt.Println(idx)
	////
	////var x,y int
	////var f uint8
	////for x=0;x <6;x+=2{
	////	y = idx[x]
	////	z:= idx[x+1]
	////	f = workArray[y]
	////	workArray[y]= workArray[z]
	////	workArray[z]= f
	////}
	////for g:=0;g

	workArray := [10]uint8{99, 151, 137, 71, 117, 187, 20, 93, 187, 67}
	// workArray := [10]uint8{}
	idx1 := [2]int{1, 2}
	idx2 := [2]int{3, 5}
	idx3 := [2]int{7, 8}

	//for i :=0; i < len(workArray); i++ {
	//	fmt.Scan(&workArray[i])
	//}

	f := workArray

	//for k := 0; k < len(idx1); k++{
	//	fmt.Scan(&idx1[k])
	//}

	f[idx1[0]], f[idx1[1]] = f[idx1[1]], f[idx1[0]]

	// ---------------------

	//for k := 0; k < len(idx2); k++{
	//	fmt.Scan(&idx2[k])
	//}

	f[idx2[0]], f[idx2[1]] = f[idx2[1]], f[idx2[0]]

	// --------------------------

	//
	//for k := 0; k < len(idx3); k++{
	//	fmt.Scan(&idx3[k])
	//}

	f[idx3[0]], f[idx3[1]] = f[idx3[1]], f[idx3[0]]

	//result := fmt.Sprintf(f)
	//fmt.Printf("%v, ==>", result)

	//fmt.Printf(fmt.Sprintf("%v", f))  // string

	for i := 0; i < len(f); i++ {
		fmt.Print(f[i], " ")
	}

}
