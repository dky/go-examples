package main

import "fmt"

func main() {
	max := 24000
	guess := 6000

	i := 0

	for i < max {
		i = i + 1
		if (i == guess) {
			fmt.Println("Found it")
			break
		}
		fmt.Println(i)
	}
}
