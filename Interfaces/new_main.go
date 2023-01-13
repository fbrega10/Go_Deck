package main

import (
	"fmt"
)

type Car struct{
	model string
	Engine
	color string
}

type Engine struct{
	horsePower int32
}

type Vehicle interface{
	toString() 
	updateVehicleTarg()
}

func (c Car) toString() {
	fmt.Println("car model : ", c.model)
}


func (c *Car) updateVehicleTarg() {
	c.model = "Fiat"	
}

func (e Engine) toString(){
	fmt.Println("engine horsepower : ", e.horsePower)
}

func (e *Engine) updateVehicleTarg(){
	e.horsePower = 117
}

// func main(){
// 	c := Car{model: "Tesla",
// 			 Engine : Engine{
// 				horsePower: 450,
// 			 },
// 			 color: "blue",
// 				}

// 	c.toString()

// 	v := Engine{
// 		horsePower: 700,
// 	}
// 	v.toString()

// 	fmt.Println("applying an update \n\n")
// 	//applying the update contract 
// 	c.updateVehicleTarg()
// 	v.updateVehicleTarg()

// 	c.toString()
// 	v.toString()
	
// }
