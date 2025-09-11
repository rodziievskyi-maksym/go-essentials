package main

import (
	"fmt"
)

func main() {
	nums := []int{9, 3, 6, 2, 5}

	selectionSort(nums)
}

func selectionSort(nums []int) {
	newArr := make([]int, len(nums)) //memery allocation  ---- [0,0,0,0,0] -> append
	length := len(nums)

	for i := 0; i < length; i++ {
		smallestIndex := findSmallest(nums) //3
		newArr[i] = nums[smallestIndex]

		//temp3 := nums                                   // {9, 3, 6, [2], 5}
		//nums = nums[:smallestIndex]                     // {9, 3, 6,}
		//nums = append(nums, temp3[smallestIndex+1:]...) // {9, 3, 6, 5}

		nums = append(nums[:smallestIndex], nums[smallestIndex+1:]...) // {9, 3, 6,} +  {5}

		// short version of code above - скорочені версія кода вище
		//newArr[i] = findSmallest(nums)
	}

	fmt.Println(newArr)
}

func findSmallest(nums []int) int {
	smallest := nums[0] // 2 ------ {9, 3, 6, 2, 5}
	smallestIndex := 0  //

	for i, num := range nums {
		if num < smallest {
			smallest = num
			smallestIndex = i // 3
		}
	}

	return smallestIndex
}
