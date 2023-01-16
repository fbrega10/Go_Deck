package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fileName := os.Args[1]
	if os.Args[1] == ""{
		fmt.Println("invalid file name: must not be empty, specify a file name")
		os.Exit(1)
	}

	source, err := os.Open(fileName)
	if err != nil{
		fmt.Println("error opening the file named : ", fileName)
	}

	io.Copy(os.Stdout, source)
}
