package main

import (
	"fmt"
	"os"
	"strconv"
)

type DayOfTheWeek uint64

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
	OtherDays
	SecondsOtherDay
)

func main() {
	args := os.Args

	numA, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}

	numB, err := strconv.Atoi(args[2])
	if err != nil {
		return
	}

	sum := numA + numB

	fmt.Printf("%d + %d = %d\n", numA, numB, sum)
}
