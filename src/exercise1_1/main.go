package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print(os.Args[0] + ": ")
	fmt.Println(strings.Join(os.Args[1:], " "))
}
