package main

import "fmt"

func main() {
	var number int

	go func() {
		number++
	}()

	if number == 0 {
		fmt.Println("number is zero: ", number)
	} else {
		fmt.Println("number is more than zero: ", number)
	}
}
