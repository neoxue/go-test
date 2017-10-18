package main

import "fmt"

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr {
		"lookback": {127,1,1,1},
		"googleDNS": {8,8,8,8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
