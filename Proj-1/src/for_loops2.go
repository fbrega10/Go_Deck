package main

import(
	"fmt"
)

// var ix = 0
// func main(){

// 	ix++

// 	for {
// 		fmt.Printf("value %d\n", ix)
// 		ix++
// 		if (ix) == 100000000{
// 			break
// 		} 
// 	}
// 	println("This is the end : reached 10 M print + conditions + value++")
// }

func main(){
	for i := 0; i < 100000000; i++{
		fmt.Printf("Current value : %d\n", i)
	}
}

