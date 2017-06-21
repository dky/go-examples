package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a) //0xc42000e268

	var b *int = &a
	fmt.Println(b) //0xc42000e268
	fmt.Println(*b) //43

	*b = 42 // set the value at memory address to 42
	fmt.Println(*b) // dereference and print out the value at b
}
