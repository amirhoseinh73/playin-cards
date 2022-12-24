package main

import (
	"os"
	"testing"
)

func TestListOfCards(t *testing.T) {
	cards := listOfCards()

	if len(cards) != 52 {
		t.Errorf("Expected cards lenght of 52, but got: %d", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got: %s", cards[0])
	}

	if cards[len(cards)-1] != "King of Diamonds" {
		t.Errorf("Expected King of Diamonds, but got %s", cards[len(cards)-1])
	}
}

var testingFileName string = "_decktesting"

func TestSaveToFileAndListOfCardsFromFile(t *testing.T) {

	os.Remove(testingFileName)

	cards := listOfCards()
	cards.saveToFile(testingFileName)

	loadedCards := listOfCardsFromFile(testingFileName)

	if len(loadedCards) != 52 {
		t.Errorf("Expected cards lenght of 52, but got: %d", len(cards))
	}

	os.Remove(testingFileName)
}
