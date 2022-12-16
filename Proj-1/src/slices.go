package main

import (
	"fmt"
)


func main(){

	newslice := make([]int, 0)
	var i = 57
	newslice = append(newslice, i)
	newslice = append(newslice, 45)
	fmt.Printf("%d \n", newslice[0])
	fmt.Printf("%d \n", newslice[1])
	for i, v := range newslice{
		fmt.Printf("index : %d, value : %d\n", i, v)
	}
}
