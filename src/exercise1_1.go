package main

import (
    "strings"
    "os"
    "fmt"
)

func main() {
    fmt.Print(os.Args[0] + ": ")
    fmt.Println(strings.Join(os.Args[1:], " "))
}
