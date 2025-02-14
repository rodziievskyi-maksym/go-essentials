package main

import "fmt"

func main() {
	//task 4

	//Stack allocation and it stores data in memory cell.
	var orange Orange

	fmt.Printf("orange memeory address %p", &orange)

	// Since we have method that will modify the struct Orange
	// I would use here an explicit pointer to Orange struct just to avoid confusion

	//newOrange holds pointer to Orange struct that allocated on the Heap
	newOrange := &Orange{}
	newOrange.Increase(10)
	newOrange.Decrease(5)
	fmt.Println(newOrange.String())

	fmt.Printf("newOrange pointer address: %p (stack)\n", &newOrange)
	fmt.Printf("newOrange memory address (heap): %p (heap)\n", newOrange)
}

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	if o.Quantity < n {
		o.Quantity = 0
		return
	}

	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("%v", o.Quantity)
}
