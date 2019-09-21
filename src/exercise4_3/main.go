package main

import "fmt"

func reverse(ptr *[7]string) {
	n := len(*ptr)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		(*ptr)[i], (*ptr)[j] = (*ptr)[j], (*ptr)[i]
	}
}

func main() {
	week := [7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
	reverse(&week)
	fmt.Printf("%v\n", week)
}
