package main

import "fmt"

//Criano um novo tipoe de deck
// que nada mais é do cque um slice de string
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
