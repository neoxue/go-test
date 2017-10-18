package main

import (
	"math"
	"fmt"
)

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}