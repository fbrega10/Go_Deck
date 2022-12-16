package main

import (
	"fmt";
	"os";
	"strconv";
)

func main(){

	if len(os.Args) < 1{
		fmt.Println("Invalid array length")
		return
	}

	y, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil{
		return
	}

	if y == 0{
		fmt.Printf("There's an error, you cannot put 0\n")
		return 
	}
	var i float64 = 0
	// i++
	for i = 0; i < y; i++{
		fmt.Printf("Counting right now: (index: %f)\n", i)
	}
	
	
}