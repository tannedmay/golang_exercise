package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	buf := &bytes.Buffer{}
	n := len(s)
	left := n % 3
	if left == 0 {
		left = 3
	}
	buf.WriteString(s[:left])
	for i := left; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}
