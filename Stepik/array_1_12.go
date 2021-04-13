package main

import "fmt"

func main() {

	//	workArray := [10]uint8{99,151,137,71,117,187,20,93,187,67}
	workArray := [10]uint8{}
	idx1 := [2]int{}
	idx2 := [2]int{}
	idx3 := [2]int{}

	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}

	f := workArray

	for k := 0; k < len(idx1); k++ {
		fmt.Scan(&idx1[k])
	}

	f[idx1[0]], f[idx1[1]] = f[idx1[1]], f[idx1[0]]

	// ---------------------

	for k := 0; k < len(idx2); k++ {
		fmt.Scan(&idx2[k])
	}

	f[idx2[0]], f[idx2[1]] = f[idx2[1]], f[idx2[0]]

	// --------------------------

	for k := 0; k < len(idx3); k++ {
		fmt.Scan(&idx3[k])
	}

	f[idx3[0]], f[idx3[1]] = f[idx3[1]], f[idx3[0]]

	for i := 0; i < len(f); i++ {
		fmt.Print(f[i], " ")
	}

}
