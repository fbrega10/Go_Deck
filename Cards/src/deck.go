package main
import (
		"fmt";
		"strings";
		"os"
	)

//typing a new type : deck 	
type deck[]string

func newDeck() deck {
	//we create a deck of cards by using for loops
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}
	
	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, suit + " of " + value)
		}
	}
	return cards
}


func (d deck) print(){
	for _, v := range d{
		fmt.Println(v)
	}
}

//this is a function which gets a deck and an int,
//returns two decks : one from zero to 
func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

//function that converts a deck (slice of strings) to a string
func (d deck) toString() string{
	return strings.Join([]string(d), ",")
	// var newString string
	// for _, value := range myarray{
	// 	newString += value
	// }
}

//function to save a deck (converted into a slice of bytes) to the hard drive
func (d deck) saveToFile(filename string) {
	os.WriteFile(filename, []byte(d.toString()), 0777)
}