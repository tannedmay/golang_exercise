package main

import (
	"bytes"
	"fmt"
	"os"
)

func commaMandisa(s string, buf *bytes.Buffer) string {
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

func comma(s string) string {
	buf := &bytes.Buffer{}
	mantissaStart := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		mantissaStart++
	}
	i := 0
	for i = mantissaStart; i < len(s); i++ {
		if s[i] == '.' {
			break
		}
	}
	mantissaEnd := i
	mantissa := s[mantissaStart:mantissaEnd]
	commaMandisa(mantissa, buf)
	buf.WriteString(s[mantissaEnd:])
	return buf.String()
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}
