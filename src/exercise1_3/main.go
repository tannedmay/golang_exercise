package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	timed(v1)
	timed(v2)
	timed(v3)
}

func timed(a func()) {
	start := time.Now()
	a()
	fmt.Println("Time taken: ", time.Now().Sub(start))
}

func v1() {
	var sep, s string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func v2() {
	sep, s := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func v3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
