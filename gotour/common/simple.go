package common

import "math"

func Add (x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x;
}

func Split(sum int) (x, y int) {
	x = sum * 4/9;
	y = sum -x;
	return
}

type Vertex struct {
	X,Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v Vertex) Scale(f float64) {
	v.X = v.X *f
	v.Y = v.Y *f
}



