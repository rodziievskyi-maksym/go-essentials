package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

type stringStruct struct {
	str unsafe.Pointer
	len int
}

func main() {
	myNameUKR := "Максим"
	myNameENG := "Maksym"
	mySurnameUKR := "Родзієвський"
	mySurnameENG := "Rodziievskyi"

	/*
		Because strings are immutable, new memory is allocated on every +=.
		To avoid lots of allocations you can opt for the more efficient strings.Builder or even bytes.Buffer
		because they have internal buffers that optimize memory allocations,
		but do that only when you deal with many concatenations, usually += is the perfect choice.
	*/
	var builder strings.Builder
	//var buffer bytes.Buffer

	//print english name address, type and size
	fmt.Printf("English name address %p, type %T, size (bytes) %d, unicode char %d\n",
		&myNameENG,
		myNameENG,
		len(myNameENG),
		utf8.RuneCountInString(myNameENG),
	)

	fmt.Printf("Ukr name address %p, type %T, size (bytes) %d, unicode char %d\n",
		&myNameUKR,
		myNameUKR,
		len(myNameUKR),
		utf8.RuneCountInString(myNameUKR),
	)

	if len(myNameENG) > len(myNameUKR) {
		fmt.Println("English name is longer")
	} else {
		fmt.Println("Ukrainian name is longer")
	}

	if utf8.RuneCountInString(myNameENG) == utf8.RuneCountInString(myNameUKR) {
		fmt.Println("Names have the same number of characters")
	}

	//strings are immutable I can't change them
	builder.Grow(len(myNameENG) + len(mySurnameENG) + 1)

	builder.WriteString(myNameENG)
	builder.WriteByte(' ')
	builder.WriteString(mySurnameENG)

	fmt.Println(builder.String())

	for i := range utf8.RuneCountInString(myNameUKR) {
		fmt.Printf("%x ", mySurnameUKR[i])
	}

	fmt.Printf("\n%x ", mySurnameUKR[1])
	fmt.Printf("\n%c ", mySurnameUKR[1])
	fmt.Printf("\n%c ", mySurnameENG[1])

	for _, v := range myNameENG {
		fmt.Print(v, " ")
	}
}
