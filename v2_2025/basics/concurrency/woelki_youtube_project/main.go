package main

import "fmt"

type Order struct {
	ID     int
	Status string
}

func main() {
	_ = generateOrders(20)

	fmt.Println("All operations complete. Exiting.")
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
