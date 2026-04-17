package main

import "fmt"

// create a new type of deck (slice of strings)
type deck []string

// function to print the deck of cards, the receiver is of type deck (d deck), so we can call this function on any variable of type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}