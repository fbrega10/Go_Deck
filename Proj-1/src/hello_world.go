package main

import ("fmt";
		"time"
		)

func main()  {
	// summ := 0
	// for i := 0; i < 11; i++{
	// 	summ += i
	// }

	// fmt.Println("The sum of the first 10 numbers is : ", summ)
	// fmt.Println("Hello world")
	// current_time := time.Now()
	
	// fmt.Println("the time is : ", current_time)

	start := time.Now()
	for i:= 0; i < 1000001; i++{
		// fmt.Println(i)
	} 
	end := time.Since(start)
	fmt.Println("benchmark : ", end)
}