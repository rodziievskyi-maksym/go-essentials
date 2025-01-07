package main

import (
	"fmt"
	"math"
)

type emptyInterface2 interface{}

func main() {
	//is there any difference?

	var emptyInterface interface{}
	var ei2 emptyInterface2
	var anyInterface any

	fmt.Printf("(%v, %T)\n", emptyInterface, emptyInterface)
	fmt.Printf("(%v, %T)\n", ei2, ei2)
	fmt.Printf("Any type it is an alias to interface{} (%v, %T)\n", anyInterface, anyInterface)

	//One of the purpose of empty interface is to handle unknown values
	handleUnknownType("type string")
}

func handleUnknownType(value interface{}) {
	//to trigger deferred anonymous recover function it has to be on top of the function stack
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("I just recover your program after a panic!: ", err)
		}
	}()
	//but if I need to know the type and check it
	//type assertion kicks in

	/*
		step 1
		we assert that empty interface value holds string type, if true we'll assign string to variable
	*/
	assertedValue := value.(string)
	/*
		if an interface doesn't hold the asserted value the statement triggers panic!
	*/
	//secondAssertion := value.(int)
	/*
		to avoid panic! triggering test assertion by receiving second value
	*/
	thirdAssertion, ok := value.(string)
	if !ok {
		fmt.Println("error occurs while asserting empty interface")
	}

	fmt.Printf("(%v, %T) can't use len() on value \n", value, value)
	fmt.Printf("(%v, %T) len:%d \n", assertedValue, assertedValue, len(assertedValue))
	fmt.Printf("Third assertion (%v, %T) len:%d \n", thirdAssertion, thirdAssertion, len(thirdAssertion))

	//fmt.Printf("after the panic here shoul be empty value of asserted type (int) %T", secondAssertion)

	/*
		Type Switches this is a construct that permits several type assertions
	*/
	var PI float32 = math.Pi
	handleUnknownTypeAnyWithTypeSwitches(PI)
}

func handleUnknownTypeAnyWithTypeSwitches(value any) {
	switch assertedValue := value.(type) {
	case int:
		fmt.Printf("(%v, %T)\n", assertedValue, assertedValue)
	case float32:
		fmt.Printf("(%v, %T)\n", assertedValue, assertedValue)
	case string:
		fmt.Printf("(%v, %T)\n", assertedValue, assertedValue)
	default:
		fmt.Println("No match!")
	}
}
