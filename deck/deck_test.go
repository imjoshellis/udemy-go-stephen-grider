package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLen := 16
	if len(d) != expectedLen {
		t.Errorf("Expected deck length of "+fmt.Sprint(expectedLen)+" but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[15] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs but got %v", d[15])
	}
}
