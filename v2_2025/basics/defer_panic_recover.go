package main

import "fmt"

func main() {
	devConfigNum := loadDevConfig(10)
	fmt.Println(devConfigNum)

	prodConfigNum := loadProdConfig(1)
	fmt.Println(prodConfigNum)
}

func loadDevConfig(num int) int {
	defer panicRecover()

	if num > 5 {
		panic("we don't have that many dev configs")
	}

	return 1
}

func loadProdConfig(num int) int {
	defer panicRecover()

	if num > 5 {
		panic("we don't have that many prod configs")
	}

	return 1
}

func panicRecover() {
	if err := recover(); err != nil {
		fmt.Println("recovered from panic:", err)
	}
}
