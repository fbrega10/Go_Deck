package main 

import (
	"os";
	"fmt";
	"strconv"
)

func main(){
	if len(os.Args) < 2{
		fmt.Println("Please provide a command line argument")
	}
	argument := os.Args[1]
	value, err := strconv.Atoi(argument)
	if err != nil{
		fmt.Println("Cannot convert to int : ", argument)
		return
	}
	
	switch {
	case value == 0:
		fmt.Println("Zero")
	case  value < 0:
		fmt.Println("Negative Value")
	case value == 100:
		fmt.Println("100!!!")
	case value > 0:
		fmt.Println("Positive Value")
		fallthrough
	default:
		fmt.Printf("Value : %d\n", value)
	}
}