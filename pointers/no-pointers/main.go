package main

import "fmt"

func zero(x int) {
	x = 0
	fmt.Println(x)
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) //You would assume that x is now 0, but it remains 5, you need to pass a pointer
}
