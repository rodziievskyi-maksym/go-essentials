package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
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
	desiredNum := 7
	indexOfDesiredNum := findIndex(oddNumber, desiredNum)
	if indexOfDesiredNum == -1 {
		log.Fatal("No index found")
	}
	fmt.Println(oddNumber[indexOfDesiredNum])

	//findIndex func already exists in slices standard library
	slices.Index(oddNumber, desiredNum)

	//generate random
	numbers := make([]float64, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = rand.Float64()
	}

	sumResult := calculateSum(numbers)
	fmt.Printf("type %T, result: %f \n", sumResult, sumResult)
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

type Number interface {
	int | int32 | int64 | float32 | float64
}

func calculateSum[T Number](numbers []T) T {
	var sum T
	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}

/*
	at first glance this type parameter confuses me S ~[]E,
	but it means (~) any type that underlay slice
*/

func SlicesIndex[S ~[]E, E comparable](slice S, value E) int {
	for i := range slice {
		if slice[i] == value {
			return i
		}
	}

	return -1
}
