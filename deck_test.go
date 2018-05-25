package main

import "testing"

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
