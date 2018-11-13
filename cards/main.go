package main

import "fmt"

func main() {
	cards := newDeck()

	// fmt.Println(cards.toString())

	// fmt.Println("****************")

	// fmt.Println(toByteSlice(cards.toString()))
	// fmt.Println(cards.toByteSlice(cards.toString()))

	// hands, remainingCards := deal(cards, 5)

	// fmt.Println("All Cards")
	// cards.print()

	// fmt.Println("Hands")
	// hands.print()
	
	// fmt.Println("Remaining Cards")
	// remainingCards.print()

	cards.saveToFile("my-cards")

	// name := ""
	// fmt.Print("Enter your name: ")
	// fmt.Scan(&name)
	// fmt.Println("Your name is", name)

}