package basics

import "fmt"

func Closures() {
	simpleClosure()
	changeValueClosure()
	pastClosureToAFunction(sumValues, 5, 7)
}

func simpleClosure() {
	intSum := func(x, y int) int { return x + y }
	result := intSum(3, 5)

	fmt.Println(result)
}

// Closures CAN CHANGE value outside a function

func changeValueClosure() {
	sample := 1
	fmt.Println("Value before closure function:", sample)

	changeVar := func() {
		sample += 1
	}

	changeVar()
	fmt.Println("Value after closure function", sample)
}

func pastClosureToAFunction(f func(int, int) int, x, y int) {
	// Here we can pass values that provided with parameters
	fmt.Println("Answer: ", f(x, y))
}

func sumValues(x, y int) int {
	return x + y
}

// Recursion Section

func Recursion() {

}
