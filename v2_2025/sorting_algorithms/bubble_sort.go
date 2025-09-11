package main

import (
	"fmt"
	"slices"
)

// example of bubble sort
func main() {
	fruits := []string{"Лимон", "Банан", "Авокадо"}

	for i := 0; i < len(fruits)-1; i++ {
		for j := 0; j < len(fruits)-i-1; j++ {
			// lexicographically comparison -> it takes first laters and compares their Unicode values.
			if fruits[j] > fruits[j+1] {
				// after comparison change position of each element.
				fruits[j], fruits[j+1] = fruits[j+1], fruits[j]
			}
		}
	}

	slices.Sort(fruits)

	fmt.Println(fruits)
}
