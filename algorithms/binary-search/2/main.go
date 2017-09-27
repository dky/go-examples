package main

import (
	"fmt"
	"os"
)

func BinarySearch(search_slice []int, value int) {
	low := 0
	high := len(search_slice) - 1
	fmt.Println("match-value:", value)
	fmt.Println("initial-low:", low)
	fmt.Println("initial-high:", high)
	fmt.Println("-------------")

	for low <= high {
		median := (low + high) / 2
		fmt.Println("median:", median)
		fmt.Println("search-slice-median:", search_slice[median])

		if search_slice[median] == value {
			fmt.Println("Matched!", value)
			os.Exit(3)
		} else if (search_slice[median] < value) {
			low = median + 1
			fmt.Println("new low:", low)
			fmt.Println("current high:", high)
		} else {
			high = median - 1
			fmt.Println("current low:", low)
			fmt.Println("new high:", high)
		}
	}
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	BinarySearch(values, 2)
}
