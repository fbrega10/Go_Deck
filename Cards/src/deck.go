package main
import (
		"fmt";
		"strings";
		"os";
		"math/rand"
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

//this is a function which has a deck and an integer as params
//returns two decks : one from zero to the handsize and from the handsize on 
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
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0777)
}

func newDeckFromFile (filename string) deck{ 
	data, err := os.ReadFile(filename)
	if err != nil{
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	data2 := strings.Split(string(data), ",")
	return deck(data2)
}

func (d deck)shuffleCards() deck{
	var temp string
	rnd := int32(len(d))
	for i:=0; i < int(rnd) ; i++{
		//random number generator swap
		gen := rand.Int31n(rnd)
		if int32(i) == gen{
			continue
		}
		temp = d[gen]
		d[gen] = d[i]
		d[i] = temp
	} 
	return d
}