package main

import (
	"fmt"
	"log"
	// "os"
	// "strconv"
)


func main(){
	
	fmt.Printf("Input the max value you intend to reach : \n")
	//declaring the variable	
	var input int

	//puntactor to the variable (input int)
	_, err := fmt.Scanln(&input)
	
	if err != nil{
		fmt.Printf("error occurred\n")
		log.Fatalln(err)
		return
	}

	if input < 1{
		fmt.Printf("Error: the int must be > 1, current value : %d\n", input)
	}
    
	for i:=0; i <= int(input); i++{
		fmt.Printf("Counting...%d\n", i)
	}
	fmt.Println("DONE!! FINISHED")

}