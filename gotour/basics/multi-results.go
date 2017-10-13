package main

import (
	"comos.sina.com/gotour/common"
	"fmt"
)

func main () {
	a, b := common.Swap("hello", `world`)
	fmt.Println(a, b)
}

