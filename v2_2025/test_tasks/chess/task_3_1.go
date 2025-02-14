package main

import (
	"fmt"
	"reflect"
)

func main() {
	//task 3_1

	//How do I compare 2 structs?

	//Simple solution, if the structs are comparable
	kingPeace := comparablePiece{
		Name:     "king",
		Color:    "white",
		isNoble:  true,
		Position: "1h",
	}

	rookPeace := comparablePiece{
		Name:     "rook",
		Color:    "white",
		isNoble:  false,
		Position: "1e",
	}

	fmt.Println(kingPeace == rookPeace)

	//Non-comparable struct with the slice of pawns
	//I would use go's built-in reflect package and a function DeepEqual or custom solution like manual comparing
	whiteQueenPiece := nonComparablePiece{
		comparablePiece: comparablePiece{
			Name:     "queen",
			Color:    "white",
			isNoble:  true,
			Position: "1d",
		},
		pawns: make([]string, 8),
	}

	blackQueenPiece := nonComparablePiece{
		comparablePiece: comparablePiece{
			Name:     "queen",
			Color:    "black",
			isNoble:  true,
			Position: "1d",
		},
		pawns: make([]string, 8),
	}

	fmt.Println(reflect.DeepEqual(whiteQueenPiece, blackQueenPiece))

	//task 3_2

	//How would I compare two interfaces?

	// If the interface contains a comparable value, I would use a simple comparison (==).
	// If the interface contains a non-comparable value, I would use the reflect package with the DeepEqual function for comparison.
	// Essentially, it's the same approach as we use for structs.
}

type comparablePiece struct {
	Name     string
	Color    string
	isNoble  bool
	Position string
}

type nonComparablePiece struct {
	comparablePiece
	pawns []string
}
