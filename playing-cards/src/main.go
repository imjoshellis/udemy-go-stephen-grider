package main

import (
	"fmt"
)

func main() {
	cards := []string{"Hi", newCard()}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}