package util

import (
	"fmt"
	"reflect"
)

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

var Pl = fmt.Println

func PrintMainSeparationMessage() {
	_, _ = Pl("---------------------------")
	_, _ = Pl("Main function processing...")
	_, _ = Pl("---------------------------\n")
}

func PrintTypeInfo(value any, isString bool) {
	if isString {
		fmt.Printf("Type %s contains of Bits: %d ",
			reflect.TypeOf(value).Name(), reflect.TypeOf(value).Size())
		return
	}

	fmt.Printf("Type %s contains of Bits: %d and has range from %d to %d\n",
		reflect.TypeOf(value).Name(), reflect.TypeOf(value).Bits(),
		-(1 << (reflect.TypeOf(value).Bits() - 1)), (1<<(reflect.TypeOf(value).Bits()-1))-1)
}

func PanicRecover() {
	if err := recover(); err != nil {
		Pl(Red+"\n***Panic Recovered***"+Reset+"\tError: ", err)
	}
}
