package main

import (
	"fmt"
	"unicode/utf8"
)

func rev(r []byte) {
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func revUnicode(slice []byte) {
	for i := 0; i < len(slice); {
		_, size := utf8.DecodeRune(slice[i:])
		rev(slice[i : i+size])
		i += size
	}
}

func main() {
	slice := []byte("Hello,  world Ø")
	revUnicode(slice)
	fmt.Println(string(slice))
}
