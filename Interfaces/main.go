package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

//very custom logic for getting an english greeting
func (englishBot) getGreeting() string{
	return "hello!"
}

func (spanishBot) getGreeting() string{
	//very custom logic for getting a spanish greeting
	return "hola!"
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

func main(){
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
