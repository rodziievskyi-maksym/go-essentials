package main

import (
	"fmt"
	"slices"
)

func main() {
	//s := []int{73, 98, 86, 61, 96}
	//fmt.Printf("slice S %v -> len %d, cap %d\n", s, len(s), cap(s))

	//s2 := s[0:2:2]
	//fmt.Printf("slice S2 %v -> len %d, cap %d\n", s2, len(s2), cap(s2))

	//s2 = append(s2, 99)

	//fmt.Printf("slice S %v -> len %d, cap %d\n", s, len(s), cap(s))
	//fmt.Printf("slice S2 %v -> len %d, cap %d\n", s2, len(s2), cap(s2))

	//fmt.Printf("S slice: %p\n", &s)
	//fmt.Printf("S2 slice: %p\n", &s2)

	//fmt.Printf("Адреса даних s:   %p\n", &s[0])
	//fmt.Printf("Адреса даних s2:  %p\n", &s2[0])

	//SparseInitializationExample()

	nums := []int{10, 11, 12, 13, 14, 10}
	containsDuplicate(nums)

	ArrayToSliceConversion()
}

func SparseInitializationExample() {
	arr1 := []int{1: 10, 2: 20}
	fmt.Println(arr1)
}

func containsDuplicate(nums []int) bool {
	slices.Sort(nums)
	length := len(nums)

	for _, num := range nums {
		if num == nums[length-1] {
			return true
		}
	}

	return false
}

func SliceToArrayConversion() {
	slice := []int{1, 2, 3, 4, 5}
	array := [5]int(slice)
	fmt.Println(array)
}

func ArrayToSliceConversion() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[2:4]

	fmt.Printf("Адреса даних s:   %p\n", &array)
	fmt.Printf("Адреса даних s2:  %p\n", &slice[0])

	fmt.Println(slice)
}
