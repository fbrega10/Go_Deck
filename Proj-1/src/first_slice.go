package main

import (
	"fmt"
)


func main(){

	aSlice := []int{-1, 2, 3, -4, 5}

	for k, v := range aSlice{
		fmt.Printf("index : %d, value : %d\n", k, v)
	}

	fmt.Println("leaving some space...")
}


