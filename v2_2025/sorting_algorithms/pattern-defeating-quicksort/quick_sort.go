package main

import "fmt"

func quickSort(arr []int, low, high int) []int {
	if low < high { // перевірка на порожність
		var p int //zero value int  = 0 //pivot - опорний елемент
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)  //Одна части
		arr = quickSort(arr, p+1, high) // Друга частина
	}

	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high] //дай нам останній елемент по індексу
	i := low           // 0

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSortStart() {
	arr := []int{5, 6, 7, 2, 1, 0}

	fmt.Println(quickSort(arr, 0, len(arr)-1)) //
}
