package main

import (
	"fmt"
)

// Converts board indices to algebraic notation (e.g., 6,4 -> "e2")
func indexToAlgebraic(rank, file int) string {
	return fmt.Sprintf("%c%d", 'a'+file, 8-rank)
}

