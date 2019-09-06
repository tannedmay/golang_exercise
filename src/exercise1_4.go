package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filesIn := make(map[string][]string)
	if len(os.Args) == 1 {
		countLines(os.Stdin, counts, filesIn)
	} else {
		command := os.Args[0]
		files := os.Args[1:]
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\t%v\n", command, err)
			}
			countLines(f, counts, filesIn)
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%20s\t%d\t%s\n", line, count, strings.Join(filesIn[line], ", "))
		}
	}
}

func in(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, filesIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if !in(filesIn[text], f.Name()) {
			filesIn[text] = append(filesIn[text], f.Name())
		}
	}
}
