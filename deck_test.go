package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 2, but got")
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 20 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {

}

func TestToString(t *testing.T) {

}
func TestToStringByInbuild(t *testing.T) {

}
func TestSaveToFilesaveToFile(t *testing.T) {

}
func TestNewDeckFromFile(t *testing.T) {

}

func TestShuffle(t *testing.T) {

}
