package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	sep = " "
	for i := 0; i < len(os.Args[1:]); i++ {
		s += os.Args[i]
		s += sep
	}
	fmt.Println(s)
}
