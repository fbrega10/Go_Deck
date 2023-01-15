package main

import (
	// "fmt"
	"fmt"
	"io"
	// "io"
	"os"

	// "io"
	"log"
	"net/http"
)

type logWriter struct{}



func main(){
	response, err := http.Get("http://www.google.com")

	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}

	// bs := make([]byte, 999999)
	// response.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	// io.Copy(os.Stdout, response.Body)
	io.Copy(lw ,response.Body)


}

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes : ", len(bs))
	return len(bs), nil
}