package basics

import (
	"fmt"
)

func MapsAndGenericsAndStructs() {
	/* This Section about Maps: */

	basicThingWithMaps()
	customTypeInMap()

	/*
		This Section about Generics
	*/
	basicThingAboutGenerics()

	/*
		This Section about Structs:
	*/
	customerFillerFunc()
	rectangleCallFunc()

	/*
		This Section about Composition
	*/

	composition()

	/*
		This Section about Define Types
	*/
	defineTypesMeasurements()
}

/* This Section about Maps: */
func basicThingWithMaps() {
	var heroes map[string]string //  one way to init map
	heroes = make(map[string]string)

	villains := make(map[string]string) // second way  to init map

	//add to map
	heroes["Batman"] = "Bruce Wayne"
	villains["Lex Luther"] = "Lex Luther"

	//another way to add
	superPets := map[int]string{
		1: "Krypto",
		2: "MeowPower",
	}

	//factor code for maps
	fmt.Printf("Batman is %v\n", heroes["Batman"])
	for i, pet := range superPets {
		fmt.Printf("Super Pet: %d: %v\n", i, pet)
	}

	//check is value there
	_, ok := superPets[5]
	if !ok {
		pl("There nothing by this index:", 5)
	}

	// Delete by key
	delete(superPets, 2)
	for i, pet := range superPets {
		fmt.Printf("Only One Pet left after deleting: %d: %v\n", i, pet)
	}
}

func customTypeInMap() {
	type myType int8
	customer := make(map[myType]string)

	customer[127] = "some-string"
	fmt.Printf("Value is %s and type is: %T\n", customer[127], customer)
}

/*This Section about Generics */

type MyConstraint interface {
	int64 | float64 | int32
}

func basicThingAboutGenerics() {
	// With the generics we can specify data type that we can use later
	// Multiple data type

	// maps
	ints := map[string]int64{
		"first":  23423421,
		"second": 323423423,
	}

	floats := map[string]float64{
		"first":  331.342343,
		"second": 342.2342,
	}

	ints32 := map[float32]int32{
		3.14: 23423421,
		3.15: 323423423,
	}

	fmt.Printf("Generic Sums: %v and %v and %v\n",
		SumIntsOrFloatsGeneric(ints),
		SumIntsOrFloatsGeneric(floats),
		SumIntsOrFloatsGeneric(ints32))
}

func SumIntsOrFloatsGeneric[K comparable, V MyConstraint](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}

	return sum
}

/* This Section about Structs: */

type Customer struct {
	name    string
	address string
	balance float32
}

func customerFillerFunc() {
	// struct literal ->  define new customer and add properties in the same time.
	userTom := Customer{"Tom", "New York", 234.43}

	CustomerInfo(userTom)

	var userJack Customer
	userJack.name = "Jack"
	userJack.address = "Boston"
	userJack.balance = 23.45

	CustomerBalance(userJack)
	newCustomerAddress(&userJack, "Dallas")
	CustomerInfo(userJack)
}

func CustomerInfo(c Customer) {
	pl("Customer info:")
	fmt.Printf("All information about customer: %+v\n", c)
}

func CustomerBalance(c Customer) {
	fmt.Printf("User name: %s\n", c.name)
	fmt.Printf("Balance: %.2f\n", c.balance)
}

func newCustomerAddress(c *Customer, address string) {
	c.address = address
}

type rectangle struct {
	length, height float32
}

func newRectangle(length float32, height float32) *rectangle {
	return &rectangle{length: length, height: height}
}

func (receiver rectangle) Area() float32 {
	return receiver.length * receiver.height
}

func rectangleCallFunc() {
	rect := newRectangle(12.12, 3.32)
	pl("Rect Area: ", rect.Area())
}

/*
	This Section about Composition
*/

type contact struct {
	firstName string
	lastName  string
	number    string
}

func newContact(firstName, lastName, num string) contact {
	return contact{
		firstName: firstName,
		lastName:  lastName,
		number:    num,
	}
}

type business struct {
	name    string
	address string
	contact
}

func newBusiness(name, address string, c contact) *business {
	return &business{
		name:    name,
		address: address,
		contact: c,
	}
}

func (b business) printInfo() {
	fmt.Printf("Contact at %s, is %s %s", b.name, b.firstName, b.lastName)
}

func composition() {
	newContact := newContact("Jon", "Doe", "4444-3333")
	newBusiness("ABC Bank", "234 North St", newContact).printInfo()
}

/*
	This Section about Define Types -> measurements
*/

type tableSpoon float64
type teaSpoon float64
type mileLiters float64

func defineTypesMeasurements() {
	mlInTableSpoon := tableSpoon(3) * 4.92
	fmt.Printf("Mileliters in %d table spoon: %f\n", 3, mlInTableSpoon)

	// Variant 2
	teaSpoonNum := teaSpoon(3)

	fmt.Printf("In %.f tea spoon(s) is %.2f mileliters\n ", teaSpoonNum, teaSpoonNum.toMileLiters())

}

func (teaSp teaSpoon) toMileLiters() mileLiters {
	return mileLiters(teaSp * 4.92)
}
