package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12{
		 t.Errorf("Expecting deck length of 12, but got %d instead", len(d))
	}
}
