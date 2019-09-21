package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		count[in.Text()]++
	}
	if in.Err() != nil {
		fmt.Fprintf(os.Stderr, "%s\n", in.Err())
		os.Exit(1)
	}
	fmt.Println(count)
}
