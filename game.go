package main

import (
    "fmt"
)

// Game represents the chess game.
type Game struct {
    board *Board
}

// NewGame initializes a new game.
func NewGame() *Game {
    return &Game{board: NewBoard()}
}

// PrintBoard displays the current board.
func (g *Game) PrintBoard() {
    g.board.Print()
}

// MakeMove applies the player's move to the game.
func (g *Game) MakeMove(move string) {
    g.board.MakeMove(move)
}

// SuggestMove suggests a move using LLaMA or another AI.
func (g *Game) SuggestMove() {
    move := getLlamaMove(g.board.GetBoardState())
    fmt.Println("Suggested Move:", move)
}

