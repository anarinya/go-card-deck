package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("☠️  🔥  Expected deck length of 48, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("☠️  🔥  Expected first card to be 'Ace of Spades', but got '%v'.", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("☠️  🔥  Expected last card to be 'King of Clubs', but got '%v'", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	// if test file exists, remove it
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("☠️  🔥  Expected 48 cards in deck, got '%v'", len(loadedDeck))
	}

	// remove test file
	os.Remove("_decktesting")
}
