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
	wg.Add(2)

	util.PrintMainSeparationMessage()

	orderCh := make(chan *Order, 20)

	go func() {
		defer wg.Done()
		defer close(orderCh)

		for _, order := range generateOrders(20) {
			orderCh <- order
		}

		fmt.Println("Done with generating orders")
	}()

	go processOrders(orderCh, &wg)

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
func processOrders(orderCh <-chan *Order, group *sync.WaitGroup) {
	defer group.Done()

	for order := range orderCh {
		time.Sleep(TimeDuration)
		fmt.Printf("Processing order %d\n", order.ID)
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
