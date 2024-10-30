package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	game := NewGame() // Initialize a new chess game
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Canary Chess")
	fmt.Println("Written by Hawk0120")

	game.PrintBoard()

	for {
		fmt.Print("Enter your move (e.g., 'e2e4') or 'suggest' for a suggested move: ")
		scanner.Scan()
		input := scanner.Text()

		if input == "suggest" {
			suggestedMove := game.SuggestMove()
			fmt.Printf("Suggested move: %s\n", suggestedMove)
			continue
		}

		if input == "quit" {
			fmt.Println("Thanks for playing!")
			break
		}

		if game.MakeMove(input) {
			fmt.Println("Move played:", input)
			game.PrintBoard()
		} else {
			fmt.Println("Invalid move. Please try again.")
		}
	}
}

