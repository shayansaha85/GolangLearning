package main

func main() {
	cards := deck{"AceOfDiamonds", newCard()}
	cards = append(cards, "SixOfSpades")

	cards.print()
}

func newCard() string {
	return "FiveOfDiamonds"
}
