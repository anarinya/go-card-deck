package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// create a new basic card deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// loop through all suits and card values, creating a full deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// prints out each card in a deck, on its own line
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal a hand given a handSize, place the rest of the cards in its own deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// shuffles the order of a given card deck

func (d deck) shuffle() {
	// create a temp deck
	t := d
	// create a new random number gen
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	// create randomized list of index nums
	p := r.Perm(len(t))

	// set card value based on randomized index
	for i := range d {
		d[i] = t[p[i]]
	}
}

// converts a deck to a single comma-delimited string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// writes a deck to a file with a given filename
func (d deck) saveToFile(filename string) error {
	// anyone can read or write the file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// reads a given filename and creates a deck based on the contents
// file is expected to contain a comma-delimited list of cards
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// if error exists, print it out and end the game
	if err != nil {
		fmt.Println("☠️  Error: ", err)
		os.Exit(1)
	}
	// convert slice into array of strings first, then deck
	return deck(strings.Split(string(bs), ","))
}
