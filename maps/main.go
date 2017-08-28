package main

import "fmt"

func main() {
	//create empty map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println(v1)
}
