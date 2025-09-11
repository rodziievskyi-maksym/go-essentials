package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		The built-in function make(T, args) serves a purpose different from new(T).
		It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T).
	*/

	// this makes a slice of length 10 and capacity 100
	// the slice is backed by an array of length 100
	//if I don't specify the capacity, the capacity will be equal to the length
	arraySliceWithCapacity := make([]int, 10, 100)
	arraySliceWithSameLengthAndCapacity := make([]int, 10)

	fmt.Printf("arraySliceWithCapacity: %v, len %d, cap %d\n", arraySliceWithCapacity, len(arraySliceWithCapacity), cap(arraySliceWithCapacity))
	fmt.Println(reflect.ValueOf(arraySliceWithCapacity).Kind())
	fmt.Printf("arraySliceWithSameLengthAndCapacity: %v, len %d, cap %d\n", arraySliceWithSameLengthAndCapacity, len(arraySliceWithSameLengthAndCapacity), cap(arraySliceWithSameLengthAndCapacity))
	fmt.Println(reflect.ValueOf(arraySliceWithSameLengthAndCapacity).Kind())

	// this makes a map with a capacity of 10 but in GO we can't check the capacity of a map
	myMap := make(map[string]int, 10+1)

	myMap["key"] = 1
	myMap["key2"] = 2

	fmt.Println(reflect.ValueOf(myMap).Kind())
	fmt.Printf("myMap: %v, len %d\n", myMap, len(myMap))

	//in go to create a channel we use make
	chName := make(chan string)
	fmt.Println(reflect.ValueOf(chName).Kind())
	fmt.Printf("chName type: %T\n", chName)

	//value of chName is a memory address
	fmt.Printf("chName: %v\n", chName)

	numberCh := make(chan int)
	messageCh := make(chan string)

	go channelDataExample(numberCh, messageCh)

	//receiving data from the channel
	fmt.Printf("Number: %d\n", <-numberCh)
	fmt.Printf("Message: %s\n", <-messageCh)
}

func channelDataExample(numCh chan int, messageCh chan string) {
	//some code

	numCh <- 10
	messageCh <- "Hello"
}
