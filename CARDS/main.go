package main

//var card string = "Ace of Spades"
func main() {

	cards := newDeck()
	// cards.print()

	//fmt.Println(cards.toString())

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	//cards.saveToFile("myCards")

	//cards := newDeckFromFile("myCa")

	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
