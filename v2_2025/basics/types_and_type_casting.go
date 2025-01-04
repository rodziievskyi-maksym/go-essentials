package main

import (
	"github.com/rodziievskyi-maksym/go-essentials/v2_2025/util"
	"unicode/utf8"
)

func main() {
	defer util.PanicRecover()
	util.PrintMainSeparationMessage()

	/*	types
		Go supports following built-in basic types:
			one boolean built-in boolean type: bool.
			11 built-in integer numeric types (basic integer types): int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, and uintptr.
			two built-in floating-point numeric types: float32 and float64.
			two built-in complex numeric types: complex64 and complex128.
			one built-in string type: string.
	*/

	/*
		bool -> needs only 1 bit to store true or false but GO uses 1 byte (smallest addressable unit of memory)
		int (based on 64-bit operating system)
		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias for uint8
		rune // alias for int32

		byte = 8 bits -> binary 0000 0000 (0 - 255)
		int16 = 16 bits -> binary 0000 0000 0000 0000 (-32768 - 32767) (0 - 65535)
		int32 = 32 bits -> binary 0000 0000 0000 0000 0000 0000 0000 0000 (-2147483648 - 2147483647)
		int64 = 64 bits -> binary (-9223372036854775808 - 9223372036854775807)
	*/
	var int64Type int = 1
	util.PrintTypeInfo(int64Type, false)

	/*
		string

		type _string struct {
			elements *byte // underlying bytes
			len      int   // number of bytes
		}

		byte -> alias for uint8
		rune -> alias for int32 (represents a Unicode code point)
		GO supports UTF-8 encoding, all Go source files must be UTF-8 encoding compatible.


	*/
	var stringType string = "Hello, World! Трохи української мови."
	util.PrintTypeInfo(stringType, true)
	util.Pl("\nAddress ", &stringType)
	util.Pl("Length in bytes: ", len(stringType))
	util.Pl("Actual characters length: ", utf8.RuneCountInString(stringType))

	/*
		float32 float64
		complex64 complex128
	*/

}
