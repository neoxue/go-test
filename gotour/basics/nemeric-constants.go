package main

import "fmt"

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x float64) float64 {
	return float64(x * 10 + 1)
}

func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Big))

	//fmt.Printf("%T", Big)
	fmt.Printf("%T %v", Small, Small)
}

