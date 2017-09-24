package main

import (
	"fmt"
	"os"
)

func main() {
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	binarySearch(numberList, 3)
}

func binarySearch(list []int, item int) {
	low := 0
	high := len(list) - 1
	mid := low + high / 2

	guess := list[mid]

	fmt.Printf("first pass: low: %d, high: %d, mid: %d\n", low, high, mid)
	fmt.Println("item is:", item)
	fmt.Println("guess is:", guess)

	//for _,i := range list {
	for low <= high {
		if guess == item {
			fmt.Println("found!")
			os.Exit(3)
			fmt.Println(mid)
		} else if (guess < item) {
			//too low
			low = mid + 1
			fmt.Println("guess < item, new low is:", low)
			os.Exit(3)
		} else if (guess > item) {
			//too high
			high = mid - 1
			fmt.Println("guess > item, new high is:", high)
			fmt.Println(high)
			os.Exit(3)
		}
		//fmt.Println(i)
	}
}
