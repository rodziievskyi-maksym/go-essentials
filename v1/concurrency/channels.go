package concurrency

import "fmt"

func main() {
	//buffered channel
	bufCh := make(chan int, 3)

	//unbuffered channel
	nonBufCh := make(chan int)

	nonBufCh <- 35
	val := <-nonBufCh

	fmt.Println(val, bufCh)
}
