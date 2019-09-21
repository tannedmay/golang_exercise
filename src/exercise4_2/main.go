package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	typ := flag.String("t", "SHA256", "Type SHA algorithm")
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arg := []byte(input.Text())
		switch *typ {
		case "SHA256":
			fmt.Printf("%x\n", sha256.Sum256(arg))
		case "SHA384":
			fmt.Printf("%x\n", sha512.Sum384(arg))
		case "SHA512":
			fmt.Printf("%x\n", sha512.Sum512(arg))
		default:
			fmt.Fprintf(os.Stderr, "Unknown SHA algorithm")
		}
	}
}
