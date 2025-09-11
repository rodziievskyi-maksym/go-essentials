package main

import "fmt"

func main() {

	// мапи мають групи замість бакетів з версії 1.24
	// як стоврюється хеш таблиця та як мапИ посилаються на неї
	// як розробники гоу вирішили питання з колізіями

	emptyMap := make(map[int]int)
	//emptyMap[0] = 0

	if emptyMap == nil {
		fmt.Println("emptyMap is nil")
	}

	emptyMap2 := map[int]int{}
	if emptyMap2 == nil {
		fmt.Println("emptyMap2 is nil")
	}
	//emptyMap2[0] = 0

	// this option considered as NIL map
	var emptyMap3 map[int]int
	if emptyMap3 == nil {
		fmt.Println("emptyMap3 is nil")
	}
	//emptyMap3[1] = 1

	fmt.Println(emptyMap)
	fmt.Println(emptyMap2)
	fmt.Println(emptyMap3)
}
