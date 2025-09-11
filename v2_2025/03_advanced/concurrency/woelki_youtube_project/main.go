package main

import (
	"fmt"
	"github.com/rodziievskyi-maksym/go-essentials/v2_2025/util"
	"math/rand"
	"sync"
	"time"
)

const TimeDuration time.Duration = time.Millisecond * 50

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	util.PrintMainSeparationMessage()

	orderCh := make(chan *Order, 20)
	processedOrderCh := make(chan *Order, 20)

	go func() {
		defer wg.Done()
		defer close(orderCh)

		for _, order := range generateOrders(20) {
			orderCh <- order
		}

		fmt.Println("Done with generating orders")
	}()

	go processOrders(orderCh, processedOrderCh, &wg)

	//select without default case is considered as blocking mechanism. It'll wait until any of cases are executed.
	go func() {
		defer wg.Done()
		for {
			select {
			case processedOrder, ok := <-processedOrderCh:
				if !ok {
					fmt.Println("Processing channel closed")
					return
				}
				fmt.Printf("Processed order %d with status %s\n", processedOrder.ID, processedOrder.Status)
			case <-time.After(time.Millisecond * 100):
				fmt.Println("Timeout is waiting for operations")
				return
			}
		}
	}()

	//for i := 0; i < 3; i++ {
	//	go func() {
	//		defer wg.Done()
	//		for _, order := range orders {
	//			updateOrderStatus(order)
	//		}
	//	}()
	//}

	wg.Wait()

	//reportOrderStatus(orders)

	fmt.Println(util.Blue + "\nAll operations complete. Exiting main.")
}

// update order status function
func updateOrderStatus(order *Order) {
	defer order.mu.Unlock()
	order.mu.Lock()
	time.Sleep(TimeDuration)
	order.Status = []string{
		"Processing", "Shipped", "Delivered",
	}[rand.Intn(3)]

	fmt.Printf("Update order %d status: %s\n", order.ID, order.Status)
}

// real world simulation
func processOrders(inOrderCh <-chan *Order, outOrderCh chan<- *Order, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(outOrderCh)
	}()

	for order := range inOrderCh {
		time.Sleep(TimeDuration)
		order.Status = "Processed"
		outOrderCh <- order
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)

	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "pending",
		}
	}

	return orders
}

func reportOrderStatus(orders []*Order) {
	fmt.Println("\n ---Order Status Report ---")

	for _, order := range orders {
		fmt.Printf("Order %d: %s\n", order.ID, order.Status)
	}

	fmt.Println("---------------------------------------\n")
}
