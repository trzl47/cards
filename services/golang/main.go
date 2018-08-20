package main

import "fmt"

func main() {
	gamePrompt()
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

type game struct {
	name    string
	aceHigh bool
}

func gamePrompt() int {
	games := []game{{name: "Blackjack", aceHigh: false}}
	fmt.Println("Welcome!")
	fmt.Println("Which game would you like to play?")
	for i, option := range games {
		fmt.Println(i, option.name)
	}
	fmt.Print("Choose game: ")
	var gameChoice int
	fmt.Scanln(&gameChoice)
	fmt.Print(gameChoice)
	return gameChoice
}
