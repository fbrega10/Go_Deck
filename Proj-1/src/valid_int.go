package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){
	arguments := os.Args
	if len(arguments) < 2{
		fmt.Printf("Invalid command line values: they must be one or more")
		return
	}

	var total, nInts, nFloats int
	invalid := make([]string, 0)

	for _, k := range arguments[1:]{
		// _ is a beautiful way to use a function, without giving an assignement
		//is it an integer? 
		_, err := strconv.Atoi(k)
		if err == nil{
			total++
			nInts++
			continue
		}
		//is it a float? 
		_, err = strconv.ParseFloat(k, 64)
		if err == nil{
			total++
			nFloats++
			continue
		}
		//then...
		invalid = append(invalid, k)
	}
	if len(invalid) > 0{
		fmt.Printf("All the invalid values here : \n")
		fmt.Printf("%s\n", invalid)
	}

	fmt.Printf("Floats : %d, Integers : %d, Total variables : %d\n\n\n\n", nFloats, nInts, total)
}