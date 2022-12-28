package main

import "fmt"


func main(){
	mySlice := []int{}
	for i := 0; i <= 10; i++{
		mySlice = append(mySlice, i)
	}
	// fmt.Println("Here's my slice of ints : \n",mySlice)
	for i, v := range mySlice{
		if v % 2 == 0{
			fmt.Println(i, v, " Even")
		}else{
			fmt.Println(i, v, " Odd")
		}
	}
}