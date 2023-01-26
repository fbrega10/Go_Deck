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

	go func(){
	for _, link := range links{
		checkLink(link, c)
	}
	close(c)
	}()

	for element := range c{
		fmt.Println(element)
	}
	
	elapsed := time.Since(start)
	fmt.Println("execution time : " + elapsed.String())
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "server down: " + link
		return
	}
	c <- "server up: " + link
}
