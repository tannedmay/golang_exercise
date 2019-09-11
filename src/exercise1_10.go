package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, uri := range os.Args[1:] {
		go fetch(uri, ch)
	}
	for range os.Args[1:] {
		fmt.Printf("%s\n", <-ch)
	}
	fmt.Printf("Time elapsed: %.2fs\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(uri, "http://") {
		uri = "http://" + uri
	}
	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %s", err)
		return
	}
	u := url.QueryEscape(uri)
	out, err := os.OpenFile(u+".out", os.O_RDWR|os.O_CREATE, 0755)
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
	ch <- fmt.Sprintf("Time elapsed: %.2fs in reading '%d' bytes from %q webpage", time.Since(start).Seconds(), n, uri)
}
