package main

import (
	"fmt"
	"os"
)

func main() {

	space := " "
	var s, value string
	for i := 1; i < len(os.Args); i++ {
		s += space + value

	}
	fmt.Println(s)
}
