package main

import "fmt"

type triangle struct{
	height float64
	length float64
}

func (t triangle ) getArea() (float64){
	return (t.height * t.length)/ 0.5 
}

type square struct{
	sideLength float64
}

func (s square) getArea() (float64){
	return s.sideLength * s.sideLength
}

type shape interface{
	getArea() float64
}

func  printArea(st string, s shape){
	fmt.Println(st, s.getArea())
}

func main(){
	//main.go -> trying out the new structs
	tr := triangle{
		height: 25.20,
		length: 12.50,
	}
	printArea("Triangle's area is : ",tr)
	sq := square{
		sideLength: 5,
	}	
	printArea("this square's area is : ", sq)
}