package main

import "fmt"

type Board struct {
	Pieces [8][8]string // 8x8 board represented by a 2D array
}

// Initializes the board to the standard starting position
func NewBoard() *Board {
	b := &Board{}
	b.Reset()
	return b
}

// Reset sets up the board with the initial chess position
func (b *Board) Reset() {
	// Set up pawns
	for i := 0; i < 8; i++ {
		b.Pieces[1][i] = "p" // White pawns
		b.Pieces[6][i] = "P" // Black pawns
	}

	// Set up main pieces for both players
	b.Pieces[0] = [8]string{"r", "n", "b", "q", "k", "b", "n", "r"} // White
	b.Pieces[7] = [8]string{"R", "N", "B", "Q", "K", "B", "N", "R"} // Black
}

// Print displays the board in the console
func (b *Board) Print() {
	fmt.Println("  a b c d e f g h")
	for rank := 0; rank < 8; rank++ {
		fmt.Printf("%d ", 8-rank)
		for file := 0; file < 8; file++ {
			piece := b.Pieces[rank][file]
			if piece == "" {
				fmt.Print(". ")
			} else {
				fmt.Printf("%s ", piece)
			}
		}
		fmt.Printf("%d\n", 8-rank)
	}
	fmt.Println("  a b c d e f g h")
}

