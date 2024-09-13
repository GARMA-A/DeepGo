package main

import (
	"fmt"
	"os"
	"testing"
)


func TestNewDeck(t *testing.T)  {

	cards := newDeck()
	if len(cards) != 16 {
	   t.Errorf("Expected deck length of 16, but got %v", len(cards))
	}
	if cards[0] != "Ace of Spades" {
		fmt.Println("Expected first card of Ace of Spades, but got", cards[0])
	}
	if cards[len(cards)-1] != "Four of Clubs" {
		fmt.Println("Expected last card of Four of Clubs, but got", cards[len(cards)-1])
	}
	
}


func TestsaveToDeckAndNewDeckFromFile(){
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		fmt.Println("Expected 16 cards in deck, got", len(loadedDeck))
	}
	os.Remove("_decktesting")

}




