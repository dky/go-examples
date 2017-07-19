package main

import "fmt"

func main() {
	//init two slices with a length of 3
	twoD := make([][]int, 4)
	//fmt.Println(twoD[3])
	//fmt.Println("Len of twoD:", len(twoD))
	//fmt.Println("Cap of twoD:", cap(twoD))
	for i := 0; i < 3; i++ {
		//fmt.Println(i)
		//returns 0 1 2
		innerLen := i + 1
		//fmt.Println(innerLen)
		//Returns 1 2 3
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			//fmt.Println(j)
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
