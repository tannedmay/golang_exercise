package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"exercise2_5/popcount"
)

func main() {
	var sha [2][32]byte
	for i, str := range os.Args[1:3] {
		sha[i] = sha256.Sum256([]byte(str))
	}
	fmt.Printf("Count of different bits b/w SHA256 of %q and %q is %d\n", os.Args[1], os.Args[2], ShaDiff(sha[0], sha[1]))
}

func ShaDiff(a, b [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += popcount.PopCountWithRightBit(uint64(a[i]) ^ uint64(b[i]))
	}
	return count
}
