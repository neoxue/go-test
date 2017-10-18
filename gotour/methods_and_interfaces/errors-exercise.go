package main

import "fmt"

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
	return "cannot sqrt negative number: -2"
}

func main () {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	fmt.Println(ErrNegativeSqrt(-2).Error())
}