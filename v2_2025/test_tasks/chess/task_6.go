package main

import "fmt"

func main() {
	// This declaration creates a nil slice of type int,
	// which is useful for nil checking to differentiate it from an empty slice.
	var a []int

	printInfo(a)

	// This declaration creates a non-nil empty slice,
	// with memory already allocated but with a length and capacity of 0.
	b := []int{}

	printInfo(b)

	// Both declarations are valid, but the second option is more commonly seen by me.
}

func printInfo(data []int) {
	fmt.Printf("type: %T, length: %d, capacity: %d\n", data, len(data), cap(data))

	if data == nil {
		fmt.Println("slice is nil")
	}
}
