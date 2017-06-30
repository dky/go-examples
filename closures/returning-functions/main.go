package main

import (
	"fmt"
)

//func that returns func
//What is the use case of this? When does it get applied? No idea right now?
func wrapper() func() int {
	//declare x @ outer scope
	x := 0
	//This is the func that's being returned by the wrapper function
	return func() int {
		//increment x inside inner scope
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
