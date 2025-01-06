package main

import (
	"fmt"
)
import "strings"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	buffer := strings.Builder{}
	buffer.Grow(len(ip)*2 - 1)

	for i := range ip {
		buffer.WriteString(fmt.Sprintf("%v", ip[i]))
		if i < len(ip)-1 {
			buffer.WriteByte('.')
		}
	}

	return buffer.String()
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
