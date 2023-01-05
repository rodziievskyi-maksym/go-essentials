package main

import "strings"

func  runSecondPart() {
	//variadicFunction(1,2,3,3,3,4,5,5,6,6,6,6,6)

	// when I initialize variable it goes to memory and gets address cell
	//myVariable := "go-essentials"
	//sendValueAsACopyFunc(myVariable)
	//pl("Output form manin: ",myVariable)

	//sendValueAsAPointer(&myVariable)
	//pl("Out from main, after pointer func", myVariable)

	// We can also store address
	//storePointerAddress()

	// Pointing on array
	//array := [4]int{1,2,3,4}
	//pointingArray(&array)
	//pl(array)


	//newSlice := []float32{11,2323,31,312.123}
	//variadicSlice(newSlice...)
}

func variadicSlice(receivingSlice ...float32) float32 {
	var sum float32
	sliceLength := float32(len(receivingSlice))
	for _, value := range receivingSlice {
		sum += value
	}

	return sum/sliceLength
}

func pointingArray(array *[4]int) {
	for i := 0; i < len(array); i++ {
		array[i] *= 2
	}
}

func storePointerAddress() {
	value := 10 // it takes space in memory
	pointerOnValue := &value
	pl("In pointerOnValue we store address in memory :", pointerOnValue) //0xc00001a0a8
	pl("Get value by pointer on it: ", *pointerOnValue) // 10
	*pointerOnValue += 5
	pl("New value in memory after modify: ",*pointerOnValue)
}

func sendValueAsAPointer(receiver *string) string {
	*receiver = strings.Replace(*receiver, "go", "golang", -1)
	pl("Output from pointer function:", *receiver)

	return *receiver
}

func sendValueAsACopyFunc(receiver string) string {
	// When I receiver value here it's copies value to a new address cell, and it considers as a new value
	changedReceiver := strings.Replace(receiver, "go", "golang", -1)
	pl("Output from function:",  changedReceiver)

	return changedReceiver
}

func variadicFunction(numbers ...int) {
	var sum int
	for _, item := range numbers {
		sum += item
	}

	pl("Summa of variadic function: ",sum)
}