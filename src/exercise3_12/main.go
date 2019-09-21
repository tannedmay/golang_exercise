package main

import (
	"fmt"
	"os"
)

func buildCharMap(s string) map[rune]int {
	charMap := make(map[rune]int)
	for _, c := range s {
		charMap[c]++
	}
	return charMap
}

func anagrams(s1, s2 string) bool {
	charMapS1 := buildCharMap(s1)
	charMapS2 := buildCharMap(s2)
	for c, count1 := range charMapS1 {
		count2, ok := charMapS2[c]
		if !ok {
			return false
		} else {
			if count1 != count2 {
				return false
			}
		}
	}
	for c, count2 := range charMapS2 {
		count1, ok := charMapS1[c]
		if !ok {
			return false
		} else {
			if count1 != count2 {
				return false
			}
		}
	}
	return true
}

func main() {
	if anagrams(os.Args[1], os.Args[2]) {
		fmt.Println("Strings are anagrams.")
	} else {
		fmt.Println("Strings are not anagrams.")
	}
}
