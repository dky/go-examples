package main

import "fmt"

func main() {
	n := 3
	//create a channel
	out := make(chan int)
	go multiplyByTwo(n, out)
	fmt.Println(<- out)
}

func multiplyByTwo(num int, out chan<- int) {
	result := num * 2
	//fmt.Println(result)
	//return result
	out <- result
}
