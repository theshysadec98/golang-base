package main

import (
	"fmt"
	"log"
	"strconv"
)

var test float64 = 3.14

const filename string = "test.txt"

func main() {

	card := newCard()
	fmt.Println(card)

	fmt.Println(test)

	fmt.Println(newInt())

	cardSlices := deck{"Your father are only testing", newCard()}
	cardSlices = append(cardSlices, adminShow(), strconv.Itoa(int(newInt())))
	cardSlices.print()

	cardWithNewDeck := newDeck()
	fmt.Println(deal(cardWithNewDeck, len(cardWithNewDeck)))
	fmt.Println(cardWithNewDeck.toString())

	err := cardWithNewDeck.saveToFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result := readDeckFile(filename)
	result.print()
	//removeDeckFileInLocal(filename)
}

func newCard() string {
	return "Five of Diamonds"
}

func newInt() int8 {
	return 111.0
}
