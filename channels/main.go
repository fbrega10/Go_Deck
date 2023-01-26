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
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.facebook.com",
	}

	c := make(chan string)

	go func() {
		for i:=0; i < len(links); i++ {
			checkLink(links[i], c)
		}
		close(c)
	}()

	for element := range c {
		fmt.Println(element)
	}

	elapsed := time.Since(start)
	fmt.Println("multi routine execution time : " + elapsed.String())

	//start of the single routine execution
	start = time.Now()

	for _, link := range links {
		singleRoutineCheck(link)
	}

	elapsed = time.Since(start)
	fmt.Println("Single routine execution time:", elapsed)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "server down: " + link
		return
	}
	c <- "server up: " + link
}

func singleRoutineCheck(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("server down: " + link)
		return
	}
	fmt.Println("server up: " + link)
}
