package concurrency

import (
	"fmt"
	"time"
)

func CountToCaller() {
	go countTo(5, "TO 5")
	go countTo(10, "TO 10")

	time.Sleep(2 * time.Second)
	fmt.Println("End of CountToCaller Function")
}

func countTo(n int, name string) {
	for i := 0; i < n; i++ {
		fmt.Printf("Function %s counting = %d\n", name, i)
	}
}

// Channel Section

func intReceiver(ch chan int, number int) {
	for i := 1; i <= number; i++ {
		ch <- i
	}
}

func IntReceiverCaller() {
	channel := make(chan int)
	go intReceiver(channel, 10)

	for i := 1; i <= 10; i++ {
		fmt.Println(<-channel)
	}
}
