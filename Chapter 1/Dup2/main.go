// This is Dup2 with task 1.4 completed

package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if len(counts[f.Name()]) == 0 {
			counts[f.Name()] = make(map[string]int)
		}
		counts[f.Name()][input.Text()]++
	}
	// Ignoring errors from input.Err()
}

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for fileName, fileValues := range counts {
		for line, n := range fileValues {
			if fileName == os.Stdin.Name() && n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			} else if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", fileName, n, line)
			}
		}
	}
}
