package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_deck")
	cards2 := newDeckFromFile("my_deck")
	cards2.print()
	cards2.shuffle()
	cards2.print()
}
