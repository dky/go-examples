package main

import "fmt"
import "time"

func main() {
	n := 3
	go multiplyByTwo(n)
	time.Sleep(time.Second)
}

func multiplyByTwo(num int) int {
	result := num * 2
	fmt.Println(result)
	return result
}
