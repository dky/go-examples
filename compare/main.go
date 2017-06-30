package main

import "fmt"

func compare(i1, i2 int) (bool, int) {
	if i1 > i2 {
		fmt.Println("The first number is greater than the second number")
		return false, i1 - i2
	} else if i2 > i1 {
		fmt.Println("The second number is greater than the first number")
		return false, i2 - i1
	} 
	fmt.Println("Numbers are equal")
	return true, 0
}

func main() {
	fmt.Println(compare(2, 0))
}
