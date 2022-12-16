package main

import (
	"fmt"
)

var y, z = 0, 0

func main(){
	// x := os.Args[1];
	for i:= 0; i < 10; i++{

		fmt.Printf("This is line n. %d\n", i)
		y ++
		z++
		if z == y{
			fmt.Printf("z == y, %d = %d\n", z, y)
		}
		
		fmt.Printf("This is y = %d, z = %d\n\n", y, z)
	}

}