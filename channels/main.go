package main

import (
	"fmt"
	"net/http"
)

func main() {

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

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "server down"
		return
	}
	c <- "server up: " + link
}
