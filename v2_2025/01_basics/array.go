package main

import "fmt"

const MagicNumber = 10
const MagicNumber2 = 20

func main() {

	userNames := getValuesFromDB()
	fmt.Println(userNames)
	fmt.Println(userNames[0])

	defer panicRecover()

	//var arr []int
	//arr2 := []int{1, 2, 3, 4}

	arrWithConstSize := [5]int64{1, 2, 3, 4, 5}

	fmt.Println(arrWithConstSize[4])

	arrWithSize := []int{1, 2, 3, 4, 5}

	arrWithSize = append(arrWithSize, 6, 7)

	//arrWithSize[2] = 10

	//for , range
	for i := 0; i < len(arrWithSize); i++ {
		arrWithSize[i] = arrWithSize[i] + MagicNumber
	}

	for i, value := range arrWithSize {
		arrWithSize[i] = value + MagicNumber2
	}

	AlwaysGetLastValue(arrWithSize)

	fmt.Printf("arrWithSize: %#v\n", arrWithSize)

	//arrStrWithSize := [7]string{"a", "b", "c", "d", "e"}
}

func AlwaysGetLastValue(arrWithSize []int) {
}

func panicRecover() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

//
//Синтаксис оголошення масиву
//Фіксований розмір масиву
//Ініціалізація масивів
//Практика: створити 3 різних масивів

//День 2: Робота з елементами масиву

//Доступ до елементів за індексом
//Зміна значень елементів
//Ітерація через масив (for, range)
//len() функція для масивів

//Практика: через цикли додати значення до кожного елементу масиву через константу
// Вивести останній елемент масиву в окремій функції
//func AlwaysGetLastValue(arrWithSize []int) {
//}

func getValuesFromDB() []string {
	return []string{"Max", "ALex", "Katya", "Tanya"}
}
