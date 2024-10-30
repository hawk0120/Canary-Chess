package main

import (
    "fmt"
)

func main() {
    game := NewGame() // Create a new game
    game.PrintBoard() // Print the initial state of the board

    // Example loop for making moves
    for {
        fmt.Print("Enter your move (or 'suggest' to get a suggestion): ")
        var input string
        fmt.Scanln(&input)

        if input == "suggest" {
            game.SuggestMove() // Get a suggestion from LLaMA
        } else {
            game.MakeMove(input) // Make the player's move
            game.PrintBoard()    // Print the updated board
        }
    }
}

