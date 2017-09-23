package main

import (
	"fmt"
	"os"
)

func main() {
	numberList := []int{1, 3, 5, 7, 9, 1, 3, 5, 9}
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

	for _,i := range list {
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
			os.Exit(3)
		}
		fmt.Println(i)
	}
}
