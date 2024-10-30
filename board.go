package main

import "fmt"

// Board represents the chessboard.
type Board struct {
    pieces [8][8]*Piece
}

// NewBoard initializes a new board with pieces in starting positions.
func NewBoard() *Board {
    b := &Board{}
    b.reset()
    return b
}

// Print prints the current state of the board.
func (b *Board) Print() {
    for rank := 0; rank < 8; rank++ {
        for file := 0; file < 8; file++ {
            if b.pieces[rank][file] != nil {
                fmt.Print(b.pieces[rank][file].Symbol() + " ")
            } else {
                fmt.Print(". ")
            }
        }
        fmt.Println()
    }
}

// Reset initializes the board with the starting pieces.
func (b *Board) reset() {
    // Set up pawns
    for file := 0; file < 8; file++ {
        b.pieces[1][file] = &Piece{Color: "white", Type: "pawn"}
        b.pieces[6][file] = &Piece{Color: "black", Type: "pawn"}
    }

    // Set up major pieces
    majorPieces := []string{"rook", "knight", "bishop", "queen", "king", "bishop", "knight", "rook"}
    for i, pieceType := range majorPieces {
        b.pieces[0][i] = &Piece{Color: "white", Type: pieceType}
        b.pieces[7][i] = &Piece{Color: "black", Type: pieceType}
    }
}

// MakeMove applies the move to the board.
func (b *Board) MakeMove(move string) {
    // Parse the move input (e.g., "e2e4") and update the board
    fromFile := int(move[0]-'a')
    fromRank := 8 - int(move[1]-'0')
    toFile := int(move[2]-'a')
    toRank := 8 - int(move[3]-'0')

    piece := b.pieces[fromRank][fromFile]
    b.pieces[toRank][toFile] = piece
    b.pieces[fromRank][fromFile] = nil
}

// GetBoardState returns the board state in a simple format.
func (b *Board) GetBoardState() string {
    // Convert the board's pieces into a custom string format
    // This is a placeholder for now
    return "Board state in FEN or custom format"
}

