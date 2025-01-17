package main

import (
	"fmt"
	"github.com/rodziievskyi-maksym/go-essentials/v2_2025/util"
	"math/rand"
	"time"
)

const TimeDuration time.Duration = time.Millisecond * 50

type Order struct {
	ID     int
	Status string
}

func main() {
	util.PrintMainSeparationMessage()

	orders := generateOrders(20)

	processOrders(orders)

	updateOrderStatuses(orders)

	reportOrderStatus(orders)

	fmt.Println(util.Blue + "\nAll operations complete. Exiting main.")
}

// update order status function
func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(TimeDuration)
		order.Status = []string{
			"Processing", "Shipped", "Delivered",
		}[rand.Intn(3)]

		fmt.Printf("Update order %d status: %s\n", order.ID, order.Status)
	}
}

// real world simulation
func processOrders(orders []*Order) {
	for _, order := range orders {
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
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(util.Magenta + "\n ---Order Status Report ---")

		for _, order := range orders {
			fmt.Printf("Order %d: %s\n", order.ID, order.Status)
		}

		fmt.Println("---------------------------------------\n")
	}
}
