package main

import (
	"fmt"
	"sync"
	"time"
)

/*
fatal error: all goroutines are asleep - deadlock!
*/
func main() {
	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum: %d\n", v1.value+1+v2.value+2)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	time.Sleep(2 * time.Second)
	go printSum(&b, &a)
	wg.Wait()
}
