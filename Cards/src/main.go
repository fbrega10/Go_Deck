package main

import "fmt"


func main() {
	//start of the program
	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_new_file")
	// cards.print()	
	
	//dealing cards into two decks
	// greeting := "Hi there"
	// fmt.Println([]byte(greeting))
}
