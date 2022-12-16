package main

import(
	"fmt";
	"math"
)


var x int = 5
var yy = x * 2

func main(){

	y := x * 5
	y = int(math.Log10(2))
	c:= 3.22

	if y <= 0 {
		c = math.Abs(float64(4))
	}else {
		c = math.Abs(float64(2))
	}
		
	//log base 10 of 2
	fmt.Printf("x=%d, y=%d, yy=%d, c=%.2f.\n", x, y, yy, c)

	fmt.Println(returnString())
}


func returnString() string{
	return "my ass"
}