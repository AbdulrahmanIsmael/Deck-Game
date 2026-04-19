package main

import "fmt"

// create a new type of deck (slice of strings)
type deck []string

// function to create a new deck of cards, it returns a variable of type deck
func newDeck() deck {
	// create an empty slice of deck cards
	cards := deck{}

	// create two slices, one for the card suits and one for the card values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// loop through the cardSuits and cardValues to create a full deck of cards | using _ to ignore the index of the loop since we don't need it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// function to print the deck of cards, the receiver is of type deck (d deck), so we can call this function on any variable of type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
