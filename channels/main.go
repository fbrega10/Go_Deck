package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	start := time.Now()
	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.facebook.com",
	}

	c := make(chan string)

	for i, link := range links {
		go checkLink(link, c)
		if i == 5 {
			close(c)
		}
	}
	index := 0
	for element := range c {
		go fmt.Println(element)
		index++
		if index == 5 {
			close(c)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("execution time : " + elapsed.String())
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "server down"
		return
	}
	c <- "server up: " + link
}
