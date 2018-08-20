package main

import (
	"fmt"
	"math/rand"
	"time"
)

// create a new type of 'deck'
// which is a slice of cards
type deck []card

type card struct {
	suit    string
	suitInt int
	face    string
	val     int
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for i, suit := range cardSuits {
		for j, face := range cardValues {
			c := card{suit: suit, suitInt: i, face: face, val: j}
			cards = append(cards, c)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card.face+" of "+card.suit)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func (d deck) toString() string {
// 	return strings.Join([]string(d), ",")
// }

// func (d deck) saveToFile(filename string) error {
// 	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
// }

// func newDeckFromFile(filename string) deck {
// 	bs, err := ioutil.ReadFile(filename) //returns []byte, error
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// s := strings.Split(string(bs), ",")
// return deck(s)
// }
