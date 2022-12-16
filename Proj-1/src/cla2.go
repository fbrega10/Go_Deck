package main

import (
	"fmt";
	"os";
	"strconv"
)

func main(){
	//standard input -> arguments value
	arguments := os.Args
	
	//first field of the standard input must be > 1 (first element represents the path)
	if len(arguments) == 1{
		fmt.Println("You need to put more arguments in the standard input")
	}

	var max, min float64 
	//two elements of float 64 type, the scope of the program is to parse each element and find the max and min values
	for i:=1; i < len(arguments); i++{
		n, err := strconv.ParseFloat(os.Args[i], 64)
		if err!= nil{
			continue
		}
		if i == 1{
			min = n
			max = n
		}
		if n < min{
			min = n
		}

		if n > max{
			max = n
		}
	}
	// fmt.Printf("\nThe max number is %.2f, the min number is %.2f\n", min, max)
	fmt.Println("max : ", max)
	fmt.Println("min : ", min)
}