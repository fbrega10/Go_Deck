package main

import(
	"fmt"
)

type fbstruct struct{
	X int
	Y int
}

func main(){
	fb := fbstruct{1, 2}

	fmt.Printf("X = %d, Y = %d\n", fb.X, fb.Y)
}