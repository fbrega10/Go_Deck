package main

import (
	"fmt"
)

func main(){
	fmt.Println("Enter your name please...")
	var name string
	fmt.Scanln(&name)
	fmt.Print("This is your name : \n ", name, "\n")
}