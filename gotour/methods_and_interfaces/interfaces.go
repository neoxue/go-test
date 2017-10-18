package main

import (
	"math"
	"comos.sina.com/gotour/common"
	"fmt"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := common.MyFloat(-math.Sqrt2)
	v := common.Vertex{3, 4}
	a = f
	a = &v
	a = v
	fmt.Println(a.Abs())
}

/*
func (f MyFloat) Abs() float64 {
	if f <0 {
		return float64(-f)
	} else {
		return float64(f);
	}
}
*/