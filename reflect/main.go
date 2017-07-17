package main

import (
	"fmt"
	"reflect"
)

func main() {
	thisString := "string"
	thisInt := 8
	thisFloat := 1.72
	thisSlice := []string{"a", "b", "c", "d"}
	thisSliceInt := []int{5, 10, 15, 20}

	fmt.Println(reflect.TypeOf(thisString))
	fmt.Println(reflect.TypeOf(thisInt))
	fmt.Println(reflect.TypeOf(thisFloat))
	fmt.Println(reflect.TypeOf(thisSlice))
	fmt.Println(reflect.TypeOf(thisSliceInt))
}
