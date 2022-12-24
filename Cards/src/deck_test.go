package main

import ("testing";
		"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12{
		 t.Errorf("Expecting deck length of 12, but got %v instead", len(d))
	}
	
	//2nd test: verifying the 1st card
	if d[0] != "Ace of Spades"{
		t.Errorf("Expected 'Ace of Spades' as first card, but got %v instead", d[0])
	}

	//3rd test : looking out for the last card
	if d[len(d)-1] != "Three of Clubs"{
		t.Errorf("Expected 'Three of Clubs' as last card, but got %v instead", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T){
	//test to see whether a new deck is the same as loaded from a file (saved before)
	os.Remove("__decktesting")
	d := newDeck()
	d.saveToFile("__decktesting")

	loadedDeck := newDeckFromFile("__decktesting")
	if len(loadedDeck) != len(d){
		t.Errorf("Expected a len of %v, got a len of %v", len(d), len(loadedDeck))
	}
}