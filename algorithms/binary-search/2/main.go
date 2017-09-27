package main

import (
	"fmt"
	"os"
)
//log2n. In this case it's 16, 2 * 2 = 4 * 2 = 8 * 2 = 16
//log2n = 4
func BinarySearch(search_slice []int, value int) {
	low := 0
	high := len(search_slice) - 1
	fmt.Println("match-value:", value)
	fmt.Println("initial-low:", low)
	fmt.Println("initial-high:", high)

	for low <= high {
		median := (low + high) / 2
		//fmt.Println("median:", median)
		fmt.Println("search-slice-median:", search_slice[median])
		//8 < 15 ?
		if search_slice[median] == value {
			fmt.Println("Matched!", value)
			os.Exit(3)
		} else if (search_slice[median] < value) {
			low = median + 1
			fmt.Println("new low:", low)
			//median is 7, is this lower than our search for value of 15?
		} else {
			high = median - 1
			fmt.Println("new high:", high)
		}
	}
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	BinarySearch(values, 4)
}
