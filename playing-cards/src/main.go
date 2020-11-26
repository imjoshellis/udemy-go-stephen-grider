package main

import (
	"fmt"
)

func main() {
	cards := deck{"Hi", newCard()}
	fmt.Println(cards)
	
	for i, card := range cards {
		fmt.Println(i, card)
	} 
}

func newCard() string {
	return "Five of Diamonds"
}