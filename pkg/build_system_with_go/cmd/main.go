package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	type DayOfTheWeek uint8

	const (
		Monday DayOfTheWeek = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
		OtherDays
		SecondsOtherDay
		ThirdOtherDay
	)

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
