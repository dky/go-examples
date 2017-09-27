package main

import (
	"fmt"
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
			break
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
	//values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	values := []int{1, 3, 5, 7, 9}
	BinarySearch(values, 3)
}
