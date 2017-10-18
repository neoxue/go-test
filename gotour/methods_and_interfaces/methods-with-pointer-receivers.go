package main

import (
	"fmt"
	"comos.sina.com/gotour/common"
)

func main() {
	v := &common.Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(4)
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
}
