package main

import "fmt"

func main() {
	//variadic function

	sumAndPrint(1, 2)
	sumAndPrint(1, 4, 5, 5, 4, 5)

	nums := []int{1, 23, 424, 24242, 2}
	//If I have slice already I can simply apply func(nums...) to variadic function
	sumAndPrint(nums...)
}

type Number2 interface {
	int | int32 | int64 | float32 | float64
}

func sumAndPrint[T Number2](num ...T) {
	//within the function, the type of nums is equivalent to slice []int
	//we can treat num as slice, iterate over it with range, use len
	fmt.Println(num, " ")

	var total T
	for _, value := range num {
		total += value
	}

	fmt.Println(total)
}
