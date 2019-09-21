package main

import "fmt"

func adjacentDup(slice []int) []int {
	index := 0
	for _, num := range slice {
		if num == slice[index] {
			continue
		}
		index++
		slice[index] = num
	}
	return slice[:index+1]
}

func main() {
	slice := []int{0, 2, 2, 2, 3, 4, 5, 5, 5}
	slice = adjacentDup(slice)
	fmt.Println(slice)
}
