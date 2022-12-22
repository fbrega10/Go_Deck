package main

import "fmt"


func main() {
	//start of the program
	cards := newDeck()
	filename := "my_new_file" 
	fmt.Println(cards.toString())
	cards.saveToFile(filename)
	cards2 := newDeckFromFile(filename)
	fmt.Println("\n\n\nDeck of cards 2 : ")
	cards2.print()
	
	fmt.Println("\n\n\n\nDeck of cards 2 sorted: ")
	cards2.shuffleCards()
	cards2.print()
	//dealing cards into two decks
	// greeting := "Hi there"
	// fmt.Println([]byte(greeting))
}
