package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {

	/* alt + command + / -> creates comment block section*/
	// out of topic
	//newYearFunction()

	//askAName() // used bufio and os.Stdin
	//defineVars()
	//dataTypesAndZeroValues() // reflect.TypeOf
	//castingIntoOtherTypes()
	//ifConditionsAgeCheck() // bufio, strings, strconv, stdin
	//letsTalkAboutStrings() // strings, replace, contains, length, split...
	//whatIsRunesInGO()
	//whatAboutTime() // time package, rand package
	//mathFunctions()
	//exampleOfCodesForPrintF() // diff format codes Sprintf example
}

func exampleOfCodesForPrintF() {
	// %d : Integer - digital
	// %c : Characters
	// %f : Float
	// %b : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	returningResult := fmt.Sprintf("%9f", 3.14) // sprintf - is the same as printf also formed in a specific format and return result
	pl(returningResult)
}

func mathFunctions() {
	pl("Pow (4,2) = ", math.Pow(4, 2)) // power
	pl("Sqrt(16) = ", math.Sqrt(16))
	pl("Cbrt(8)= ", math.Cbrt(8)) // cube
	pl("Ceil(4.3)= ", math.Ceil(4.3))
	pl("Floor(4.8)= ", math.Floor(4.8))
	pl("Round(4.4)= ", math.Round(4.4)) // 4.5 becomes 5

	pl("Log2(8) = ", math.Log2(8)) // logarithms
	pl("Log(7.389) = ", math.Log(math.Exp(2)))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90) // "%.2f -> means place left 2 decimal after a dot

	pl("Sin(90) = ", math.Sin(r90))
}

func whatAboutTime() {
	now := time.Now()
	pl(now)
	pl(now.Year(), now.Month())
	pl(now.Minute(), now.Second())

	seedSeconds := time.Now().Unix() // return seconds since 1.1.1970
	pl("Seconds since 1970: ", seedSeconds)
	rand.Seed(seedSeconds)
	randomNum := rand.Intn(50) + 1
	pl(randomNum)
}

func whatIsRunesInGO() {
	// runs are character represented by unicode
	runeString := "any text"
	pl("Bytes counts", len(runeString))
	pl("Rune Count:", utf8.RuneCountInString(runeString))
	for i, runeValue := range runeString {
		fmt.Printf("%d : %#U : %c\n", i, runeValue, runeValue)
	}
	/*
		0 : U+0061 'a' : a
		1 : U+006E 'n' : n
		2 : U+0079 'y' : y
		3 : U+0020 ' ' :
		4 : U+0074 't' : t
		5 : U+0065 'e' : e
		6 : U+0078 'x' : x
		7 : U+0074 't' : t

	*/

}

func newYearFunction() {
	pl("Slava Ukraine")
	pl("Heroyam Slava")
	pl("Smert' vorogam!!!")

	howManyPidorsLeft := 100000
	pl("How many pidors left", howManyPidorsLeft)

	pidorsLeft := 100000
	for i := 0; i < howManyPidorsLeft; i++ {
		pidorsLeft = pidorsLeft - 1
		pl(pidorsLeft)
	}

	pl("Peremoga!!!!")
}

func letsTalkAboutStrings() {
	justAString := "A new word"
	replacer := strings.NewReplacer("A", "Another")
	replacedString := replacer.Replace(justAString)

	pl(replacedString)
	pl("Length:", len(replacedString))
	pl("Contains", strings.Contains(replacedString, "Another"))

	if strings.Contains(replacedString, " ") {
		spaceIndex := strings.Index(replacedString, " ")
		pl("Index", spaceIndex)
		pl("Character:", spaceIndex+1)
	}

	pl("Replace:", strings.Replace(replacedString, "o", "0", -1))

	spacedString := "\nSome Words\n"
	pl(spacedString)
	spacedString = strings.TrimSpace(spacedString)
	pl(spacedString)

	pl("Split:", strings.Split(replacedString, " "))
	pl("To lower case:", strings.ToLower(replacedString))
	pl("To Upper case", strings.ToUpper(replacedString))

}

func ifConditionsAgeCheck() {
	// Conditional Operators : > < >= <= == !=
	// Logical Operations: && || !
	pl("Enter your Age:")

	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	result = strings.TrimSpace(result)
	var userAge int
	if userAge, err = strconv.Atoi(result); err != nil {
		log.Fatal(err)
	}

	// Age Check
	if userAge >= 1 && userAge < 18 {
		pl("Such a nice age, jealous")
	} else if userAge == 18 || userAge == 21 {
		pl("You can do whatever you what now")
	} else {
		pl("My greetings to a mature person")
	}

}

func castingIntoOtherTypes() {
	floatNumber := 2.9
	convertedIntoInt := int(floatNumber) // flooring number down even 2.9 becomes 2

	pl(convertedIntoInt)

	// how to convert string into integer
	// https://pkg.go.dev/strconv
	stringInteger := "5999999"
	converted, err := strconv.Atoi(stringInteger) // ascii table to integer
	if err != nil {
		log.Fatal(err)
	}

	num := 234234
	convertedIntoString := strconv.Itoa(num)

	pl(converted, err, reflect.TypeOf(converted))
	pl(convertedIntoString, err, reflect.TypeOf(convertedIntoString))

	// string to float number
	floatString := "3.14"
	if convertedIntoFloat, err := strconv.ParseFloat(floatString, 64); err == nil {
		pl(convertedIntoFloat)
	}

	floatIntoString := fmt.Sprintf("%f", 3.14)
	pl(floatIntoString)
}

func dataTypesAndZeroValues() {
	/*
		data types and their default values:

		int -> 0
		float -> 0.0
		bool -> false
		string -> ""
		rune -> ""
	*/

	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("true"))

}

func askAName() {
	// https://pkg.go.dev/bufio@go1.19.4
	_, err := pl("What is your name?")
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	pl("Hello", name)
	if err != nil {
		return
	}
}

func defineVars() {
	// variables
	var anyNameInCamelCase string = "string name"
	var anyName = 1.2
	justAName := "any string"
	justAName = "omit the colon to reassign a variable"

	pl(anyName, anyNameInCamelCase, justAName)
}