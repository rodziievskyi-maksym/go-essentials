package graceful_shutdown

import (
	"context"
	"log"
	"time"
)

func GracefulShutdown() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("2 seconds passed, cancel context")
		cancel()
	}()

	log.Println("waiting for context cancellation")
	<-ctx.Done()
	log.Println("context done")
}
