package main

import (
	"fmt"
	"sort"
)

func main() {
	chessPiecePositionMap := make(map[string]string, 5)
	chessPiecePositionMap = map[string]string{
		"1b": "bishop",
		"1c": "knight",
		"1a": "rook",
		"1e": "king",
		"1d": "queen",
	}

	positions := make([]string, len(chessPiecePositionMap))
	index := 0
	for position := range chessPiecePositionMap {
		//explicitly avoid append function for demonstration purposes
		//slice pre-allocation more expensive than then int variable
		positions[index] = position
		index++
	}
	sort.Strings(positions)

	for _, position := range positions {
		fmt.Printf("Position: %s, Piece: %s\n", position, chessPiecePositionMap[position])
	}
}
