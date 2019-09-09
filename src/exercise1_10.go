package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Printf("%s\n", <-ch)
	}
	fmt.Printf("Time elapsed: %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %s", err)
		return
	}
	outfile := resp.Location().Host + ".out"
	fmt.Println(outfile)
	out, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		ch <- fmt.Sprintf("Error in opening out file: %s", err)
		return
	}
	n, err := io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error in reading response: %s", err)
		return
	}
	ch <- fmt.Sprintf("Time elapsed: %.2fs in reading '%d' bytes from %q webpage", time.Since(start).Seconds(), n, url)
}
