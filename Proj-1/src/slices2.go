package main

import (
	"fmt"
)


func main(){
	//newslice := make([]string 0)
	newslice := []string{}
	var i = "apple"
	//append a new slice 
	newslice = append(newslice, i)
	newslice = append(newslice, "Banana")
	
	//fmt.Printf("%s \n", newslice[0])
	//fmt.Printf("%s \n", newslice[1])
	for i, v := range newslice{
		fmt.Printf("index : %d, value : %s\n", i, v)
		}
	fmt.Printf("You reached the end of the program, goodbye :)\n")
}
