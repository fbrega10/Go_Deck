package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	arguments := os.Args
	if len(arguments) == 1{
		fmt.Printf("We need one or more input arguments")
		return
	}
	// var x int
	var y []int
	for i := 1; i < len(arguments); i++{
		n, err := strconv.ParseInt(arguments[i], 64, 64)
		if err != nil{
			continue
		}
		for t := 0; t < len(y); t++{
			if int(n) > y[t]{
				return
			}
		} 
	}

}