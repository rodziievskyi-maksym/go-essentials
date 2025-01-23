package main

import "fmt"

func main() {
	//this is nil channel Uninitialized it leads to deadlock if I communicate with it
	var ch chan int

	//Explicitly setting a channel to Nil to prevent communication
	ch2 := make(chan int)
	ch2 = nil

	//receiving and sending to nil channel leads to deadlock
	ch <- 1
	data := <-ch2

	fmt.Println(data)

	//the case with select will never execute and ignored
	var ch3 chan int
	select {
	case ch3 <- 1:
	default:
		fmt.Printf("Default case executed")
	}
}
