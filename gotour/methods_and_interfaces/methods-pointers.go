package main

import (
	"fmt"
	"comos.sina.com/gotour/common"
)



func main () {
	v := common.Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}