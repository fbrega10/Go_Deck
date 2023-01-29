package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	tech := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.facebook.com",
	}

	banking := []string{
		"https://www.intesasanpaolo.com/",
		"https://www.unicredit.it/",
		"https://www.mediolanum.it/",
		"https://www.generali.it/",
		"https://www.allianz.it/",

	}

	c1, c2 := make(chan string), make(chan string)

	for _, i := range tech{
		go checkLink(i, c1)	
	}

	for _, s := range banking{
		go checkLink(s, c2)
	}

	for i:=0; i < (len(tech)+ len(banking)); i++{
		select {
		case msg1 := <- c1:
			fmt.Println("techs got first: ", msg1)
		case msg2 := <- c2:
			fmt.Println("banking got first: ", msg2)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("program duration : ", elapsed)

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "server down:" + link
		return
	}
	c <- "server up: " + link
}