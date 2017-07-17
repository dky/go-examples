package main

import (
	"fmt"
	"os"
_	"path/filepath"
	"reflect"
)

func main() {
	//filePath, _ := filepath.Abs("./")
	f, err := os.Open("./test.txt")
	fmt.Println("Type is:", reflect.TypeOf(f))
	//err is not just a string
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	//when done close the file
	defer f.Close()

	//read 100 bytes into b slice.
	b := make([]byte, 100)
	fmt.Println("Type is:", reflect.TypeOf(b))

	//err redeclared? No it's the same var =: checks to see if err
	//has been used before and re-uses it.
	n, err := f.Read(b)
	fmt.Printf("%d: % x\n", n, b)
	//Does s automatically convert byte into string now?
	fmt.Printf("%d: % s\n", n, b)

	someString := "foo bar"
	fmt.Printf("Type is:", reflect.TypeOf(someString))

	byteString := []byte(someString)
	fmt.Printf("Type is:", reflect.TypeOf(byteString))

}
