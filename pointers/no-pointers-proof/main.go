package main

import "fmt"

func zero(x int) {
	x = 0
	fmt.Println("Value in zero for x:", x) //x returns 0
	fmt.Printf("Address in zero: %p\n", &x) // return memory address of x
}

func main() {
	x := 5
	fmt.Printf("Address in main: %p\n", &x)
	//fmt.Println(&x)
	zero(x)
	fmt.Println("Value in main for x:", x) //You would assume that x is now 0, but it remains 5, you need to pass a pointer
}
