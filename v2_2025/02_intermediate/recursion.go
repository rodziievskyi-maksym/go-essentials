package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(sumOfDigits(12345))
	fmt.Println(fibonacci(10))

	dbResponse := DBResponseImitation(10)

	fmt.Println("money left:", activateGiftCardRecursive(105, dbResponse, 0))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func activateGiftCardRecursive(amount uint64, transaction []UserGiftCard, index int) uint64 {
	if amount < 10 {
		return amount
	}

	redeemed := transaction[index].Redeemed
	index++

	return activateGiftCardRecursive(amount-redeemed, transaction, index)
}
