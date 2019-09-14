package main

import (
	"exercise2_4/popcount"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
			os.Exit(1)
		}
		count := popcount.PopCountWith64Shift(uint64(num))
		fmt.Printf("Pop count: %d in number: %d (%s)\n", count, num, strconv.FormatInt(int64(num), 2))
	}
}
