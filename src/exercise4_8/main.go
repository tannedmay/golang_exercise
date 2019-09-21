package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	count := make(map[rune]int)
	var ulen [utf8.UTFMax + 1]int
	number := 0
	letter := 0
	control := 0
	graphic := 0
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %s", err)
			os.Exit(1)
		}
		switch {
		case unicode.IsControl(r):
			control++
		case unicode.IsGraphic(r):
			graphic++
		case unicode.IsLetter(r):
			letter++
		case unicode.IsNumber(r):
			number++
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		count[r]++
		ulen[n]++
	}
	fmt.Println(count, ulen, number, letter, control, graphic, invalid)
}
