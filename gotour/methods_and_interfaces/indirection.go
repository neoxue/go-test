package main

import (
	"comos.sina.com/gotour/common"
	"fmt"
)

func main() {
	v := common.Vertex{3, 4}
	v.Scale(3)
	common.ScaleFunc(&v, 3)
	fmt.Println(v)

	p := &common.Vertex{3, 4}
	p.Scale(3)
	common.ScaleFunc(p, 3)
	fmt.Println(p)
}
