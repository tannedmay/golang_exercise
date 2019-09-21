package main

import "fmt"

func rotate(slice []int, n int) {
	first := make([]int, n)
	copy(first, slice[:n])
	copy(slice, slice[n:])
	copy(slice[len(slice)-n:], first)
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	rotate(slice, 2)
	fmt.Println(slice)
}
