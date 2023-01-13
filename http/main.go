package main

import (
	// "fmt"
	"io"
	// "io"
	"os"

	// "io"
	"log"
	"net/http"
)



func main(){
	response, err := http.Get("http://www.google.com")

	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}

	// bs := make([]byte, 999999)
	// response.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, response.Body)


}