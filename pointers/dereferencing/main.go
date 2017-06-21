package main

import "fmt"

func main() {
	
	a := 43
	fmt.Println(a)

	fmt.Println(&a) //Think of & = A for memory address - 0xc42000e268

	var b *int = &a
	fmt.Println(b) //0xc42000e268

	fmt.Println(*b) //43 - the * operator is used to dereference memory location - Give me the value at this memory address.
}
