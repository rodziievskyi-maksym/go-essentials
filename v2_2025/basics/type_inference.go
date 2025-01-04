package main

var lastName, himAge, isSingle, himGroup = "Max", 30, false, 6.1214

func main() {
	/*
		GO can inference the type of the variable based on the value assigned to it, but it is not a dynamic type language
		so once the type is assigned, it cannot be changed
	*/
	var name, age, isMarried, group = "Max", 30, false, 6.1214

	//or

	/*
		short declaration syntax works only inside the function
	*/
	name2, age2, isMarried2, group2 := "Max", 30, false, 6.1214

	println(name, age, isMarried, group, name2, age2, isMarried2, group2)
}
