package main

import (
 "strings"
)

type Game struct {
	Board *Board
	Turn  string // "white" or "black"
}

// Initialize a new game
func NewGame() *Game {
	return &Game{
		Board: NewBoard(),
		Turn:  "white",
	}
}

// MakeMove applies a move given in algebraic notation (e.g., "e2e4")
func (g *Game) MakeMove(input string) bool {
	if len(input) != 4 {
		return false
	}

	fromRank, fromFile := algebraicToIndex(input[:2])
	toRank, toFile := algebraicToIndex(input[2:])
	piece := g.Board.Pieces[fromRank][fromFile]

	// Check if it's the correct player's turn
	if (g.Turn == "white" && piece == strings.ToUpper(piece)) || (g.Turn == "black" && piece == strings.ToLower(piece)) {
		g.Board.Pieces[toRank][toFile] = piece
		g.Board.Pieces[fromRank][fromFile] = ""
		g.Turn = oppositeTurn(g.Turn) // Switch turn
		return true
	}
	return false
}

// SuggestMove suggests a random valid move for the current player
func (g *Game) SuggestMove() string {
	// Generate a random move (just for demonstration)
	return "e2e4"
}

// PrintBoard prints the board by calling the Board's print method
func (g *Game) PrintBoard() {
	g.Board.Print()
}

// Convert algebraic notation to board indices
func algebraicToIndex(algebraic string) (int, int) {
	file := int(algebraic[0] - 'a')
	rank := 8 - int(algebraic[1]-'0')
	return rank, file
}

// Get the opposite turn
func oppositeTurn(turn string) string {
	if turn == "white" {
		return "black"
	}
	return "white"
}

