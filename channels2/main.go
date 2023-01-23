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
	for _, i := range links{
		checkLink(i)	
	}
	elapsed := time.Since(start)
	fmt.Println("program duration : ", elapsed)

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("server down")
		return
	}
	fmt.Println("server up: " + link)
}