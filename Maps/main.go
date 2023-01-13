package main

import "fmt"


func main(){
	//three ways to declarate a new map

	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#bb000",
		"white": "#4bb000",
	}
	printMap(colors)
	//add a key to the map 
	// colors["white"]= "abcd"
	// fmt.Println(colors)
	// delete(colors, "white")
	// fmt.Println(colors)
}

func printMap(c map[string]string){
	for k, v := range c{
		fmt.Println("hex code for color:", k, "is", v)
	}
}