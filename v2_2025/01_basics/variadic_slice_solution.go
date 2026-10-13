package main

import "fmt"

var users []string

func addUser(user ...string) {
	users = append(users, user...)
}

func main() {
	addUser("Bob", "Mike", "Jack")
	addUser("Mike")
	addUser("Jack")

	fmt.Println(users)
}
