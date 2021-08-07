package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// Maneira alternativa de rescrever a declaração de variável acima
	card := newCard()
	// card = "five of diamonds"
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
