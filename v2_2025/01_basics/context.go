package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//context is a package of go's standard library
	//context is immutable it can't be modified by children

	//context deadline
	ctx := context.Background()
	exampleTimeout(ctx)
}

func exampleTimeout(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*4)
	defer exampleOfClosure(cancel)

	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("Request to an API")
		close(done)
	}()

	select {
	case <-done:
		fmt.Printf("API Response")
	case <-ctxWithTimeout.Done():
		fmt.Printf("Request time exedded %s:", ctxWithTimeout.Err())
	}
}

// Closure
func exampleOfClosure(cancelFunc context.CancelFunc) {
	cancelFunc()
}
