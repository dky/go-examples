package main

import (
	"fmt"
//	"os"
)

//log2n. In this case it's 16, 2 * 2 = 4 * 2 = 8 * 2 = 16
//log2n = 4

func BinarySearch(search_slice []int, value int) int {
	low := 0
	high := len(search_slice) - 1
	fmt.Println("high:", high)

	for low <= high {
		median := (low + high) / 2
		fmt.Println("med:", median)

		if search_slice[median] < value {
			low = median + 1
			fmt.Println("new low:", low)
		//median is 7, is this lower than our search for value of 15?
		} else {
			high = median - 1
			fmt.Println("new high:", high)
		}

	}

	if low == len(search_slice) || search_slice[low] != value {
		return -1
	} else {
		return low
	}

}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	fmt.Println(BinarySearch(values, 15))
}
