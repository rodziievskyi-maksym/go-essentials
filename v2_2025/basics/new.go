package main

import (
	"fmt"
	"unsafe"
)

/*
	When you define a struct in Go,
	it does not allocate memory until you create an instance of it.
*/

type User struct {
	FirstName string
	LastName  string

	isAdmin bool
	isUser  bool
}

func main() {
	/*
		The expressions new(File) and &File{} are equivalent.
	*/
	structUser := User{}
	newUser := new(User)

	fmt.Printf("structUser type: %T\n", structUser)
	fmt.Printf("structUser fields: %+v\n", structUser)
	fmt.Printf("structUser size: %d bytes\n", unsafe.Sizeof(structUser))

	fmt.Printf("newUser type: %T\n", newUser)
	fmt.Printf("newUser fileds: %+v\n", newUser)
	fmt.Printf("newUser size: %d bytes\n", unsafe.Sizeof(newUser))
}
