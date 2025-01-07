package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		Generics it is a GO function that works with multiple types using type parameters
	*/

	normalFunc()

	//usage of generic function
	oddNumber := make([]int, 5)

	var index int
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			oddNumber[index] = i
			index++
		}
	}

	fmt.Println(oddNumber)

	//what is the index of 7 in this slice
	desiredNum := 10
	indexOfDesiredNum := findIndex(oddNumber, desiredNum)
	if indexOfDesiredNum == -1 {
		log.Fatal("No index found")
	}
	fmt.Println(oddNumber[indexOfDesiredNum])
}

// normal function
func normalFunc() {
	fmt.Println("Hello from the normal function!")
}

//generic function -> [] brackets are embedded before function arguments!
/*
	What is comparable?

 	comparable is an interface that is implemented by all comparable types
 	(booleans, numbers, strings, pointers, channels, arrays of comparable types,
 	structs whose fields are all comparable types).
*/
func findIndex[T comparable](oddNumbers []T, num T) int {
	for i, number := range oddNumbers {
		if number == num {
			return i
		}
	}

	return -1
}
