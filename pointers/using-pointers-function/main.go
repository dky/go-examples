package main

import "fmt"

func zero(z *int) {
	*z = 0
	fmt.Println("z address in zero:", z)
}

func main() {
	x := 10
	zero(&x)
	fmt.Println("x address in main:", &x)
	fmt.Println(x) //x is 0 since we are passing a memory location to the func
}
