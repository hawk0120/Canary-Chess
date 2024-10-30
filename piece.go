package main

// Piece represents a chess piece.
type Piece struct {
    Color string // "white" or "black"
    Type  string // e.g., "pawn", "knight", "bishop", etc.
}

// Symbol returns the character that represents the piece on the board.
func (p *Piece) Symbol() string {
    switch p.Type {
    case "king":
        if p.Color == "white" {
            return "K"
        }
        return "k"
    case "queen":
        if p.Color == "white" {
            return "Q"
        }
        return "q"
    case "rook":
        if p.Color == "white" {
            return "R"
        }
        return "r"
    case "bishop":
        if p.Color == "white" {
            return "B"
        }
        return "b"
    case "knight":
        if p.Color == "white" {
            return "N"
        }
        return "n"
    case "pawn":
        if p.Color == "white" {
            return "P"
        }
        return "p"
    }
    return "."
}

