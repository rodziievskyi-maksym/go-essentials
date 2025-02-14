package main

import "fmt"

func main() {
	//task 2
	castling()
	copySlice()
	copyMap()
}

func castling() {
	rookPosition, kingPosition := "1e", "1h"

	//I would use Go's multiple assignment
	kingPosition, rookPosition = rookPosition, kingPosition
}

func copySlice() {
	chestBoardColumn := []int{1, 2, 3, 4, 5, 6, 7, 8}

	copiedChestBoardColumn := make([]int, len(chestBoardColumn))
	copy(copiedChestBoardColumn, chestBoardColumn)

	fmt.Printf("memory address of original board column slice: %p\n", &chestBoardColumn)
	fmt.Printf("memory address of copied board column slice: %p\n", &copiedChestBoardColumn)
}

func copyMap() {
	piecesStartingPositionMap := map[string]string{
		//I hope 3 is enough
		"rook":   "1a",
		"knight": "1b",
		"bishop": "1c",
	}

	//I know only the way of manually coping map to the other map
	copiedPositionsMap := make(map[string]string, len(piecesStartingPositionMap))
	for piece, position := range piecesStartingPositionMap {
		copiedPositionsMap[piece] = position
	}

	fmt.Printf("memory address of original positions map: %p\n", &piecesStartingPositionMap)
	fmt.Printf("memory address of copied positions map: %p\n", &copiedPositionsMap)
}
