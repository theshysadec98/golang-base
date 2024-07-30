package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Expected deck length of %v", len(d))
	}
}
