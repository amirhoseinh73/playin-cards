package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func listOfCards() deck {

	var cards deck

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func listOfCardsFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	ErrorHandled(err)

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range d {
		np := r.Intn(len(d) - 1)

		d[i], d[np] = d[np], d[i]
	}
}

func newDeck() deck {
	cards := listOfCards()
	cards.shuffle()
	return cards
}
