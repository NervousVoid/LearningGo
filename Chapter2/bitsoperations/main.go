package main

import "fmt"

var x uint8 = 1<<1 | 1<<5
var y uint8 = 1<<1 | 1<<2

func main() {
	fmt.Printf("%08b\n", x)    // множество {1, 5} "00100010"
	fmt.Printf("%08b\n", y)    // множество {1, 2} "00000110"
	fmt.Printf("%08b\n", x&y)  // пересечение {1} "00000010"
	fmt.Printf("%08b\n", x|y)  // объединение {1, 2, 5} "00100110"
	fmt.Printf("%08b\n", x^y)  // симметричная разность {2, 5} "00100100"
	fmt.Printf("%08b\n", x&^y) // разность {5} "00100000"

	for i := uint(0); i < 8; i++ {
		if x&(1<<1) != 0 { // проверка принадлежности множеству
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // множество {2, 6}
	fmt.Printf("%08b\n", y>>1) // множество {0, 4}
}
