package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

// create a new custom type called 'deck' 
// which is a type of string
type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuites := []string { "Spades", "Hearts", "Diamonds" }
	cardValues := []string { "Ace", "Two", "Three", "Four" }

	for _, suite := range cardSuites{
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suite)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}

func toByteSlice(s string) []byte {
	return []byte(s)
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile( fileName, []byte(d.toString()), 0666 )
}

                            