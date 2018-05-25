package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	requiredLength := 16
	firstCardInDeck := "Ace of Spades"
	lastCardInDeck := "Four of Clubs"

	if len(d) != requiredLength {
		t.Errorf("Expected length of %d, but got: %d", requiredLength, len(d))
	}

	if d[0] != firstCardInDeck {
		t.Errorf("Expected first item in deck is: %s, but got: %s", firstCardInDeck, d[0])
	}
	if d[len(d)-1] != lastCardInDeck {
		t.Errorf("Expected last item in deck is: %s, but got: %s", lastCardInDeck, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	requiredLength := 16

	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != requiredLength {
		t.Errorf("Expected %d, but got: %d", requiredLength, len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	deck := newDeck()

	remainingSize, handSize := deal(deck, 3)

	if remainingSize == nil {
		t.Errorf("Expected to have %v and a hand size of: %v, but got %v", remainingSize, handSize, nil)
	}
}
