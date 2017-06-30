package main

import (
	"fmt"
)

func main() {
	x := 0

	// assign anonomous function to the increment var
	// anonomous function = a function without a name
	// func expression = assigning a function to a variable
	increment := func () int {
		x++
		return x
	}
	
	fmt.Println(increment())
	fmt.Println(increment())
}
