package main

import ( 
	"os"
	"fmt"
)

func main() {
	argsWithProg := os.Args
	fmt.Println(argsWithProg)

	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	arg := os.Args[3]
	fmt.Println(arg)
}

