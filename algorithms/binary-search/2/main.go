package main

import (
	"fmt"
)

func BinarySearch(target_map []int, value int) int {

	low := 0
	high := len(target_map) - 1

	for low <= high {

		median := (low + high) / 2
		fmt.Println(median)

		if target_map[median] < value {
			low = median + 1
		} else {
			high = median - 1
		}

	}

	if low == len(target_map) || target_map[low] != value {
		return -1
	} else {
		return low
	}

}

func main() {
	//16, 8, 4, 2, 1
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	fmt.Println(BinarySearch(values, 15))
}
