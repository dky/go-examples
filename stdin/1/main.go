package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
