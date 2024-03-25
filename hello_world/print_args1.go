package main

import (
	"fmt"
	"os"
)

func main() {

	space := " "
	var s string
	for i := 1; i < len(os.Args); i++ {
		s += space + os.Args[i]

	}
	fmt.Println(s)
}
