package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a) //returns 43
	fmt.Println(&a) //returns 0xc42000e268 - memory location/address where a is stored.

	var b *int = &a //variable b is a pointer to an int, then assign memory location of a to b
	fmt.Println(b) //returns 0xc42000e268

}
