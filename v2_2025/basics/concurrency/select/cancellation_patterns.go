package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))

	defer cancel()

	//go worker(ctx)
	go workerWithCleanUp(ctx)

	// Wait for cancellation
	<-ctx.Done()
	//fmt.Println("Main: worker cancelled due to: ", ctx.Err())

	fmt.Println("Main: Waiting for cleanup...")
	time.Sleep(2 * time.Second) // Wait for cleanup to finish

}

func worker(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			//simulate work
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Working")
		}
	}
}

func workerWithCleanUp(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		// Cleanup when cancelled
		case <-ctx.Done():
			fmt.Println("Worker: Cleaning up resources...")
			time.Sleep(1 * time.Second) // Simulate cleanup
			fmt.Println("Worker: Cleanup done")
			return
		case <-ticker.C:
			fmt.Println("Worker: Doing work")
		}

	}
}
