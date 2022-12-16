package main

import (
	"fmt";
	"os";
//	"strconv"
)


func main(){
	if len(os.Args) != 2{
		fmt.Println("Please provide a command line argument")
		return
	}
	argument := os.Args[1]
	// With expression after switch 
	switch argument{
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2":
		fmt.Println("Two!")
	case "3":
		fmt.Println("Three!")
	case "4","5","6","7","8","9","10":
		fmt.Println("4 - 10")
		fallthrough //tells Go to go after the deafault value
	default:
		fmt.Println("Value : ", argument)
	}
}
