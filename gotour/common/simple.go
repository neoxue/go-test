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



type Vertex2 struct {
	X float64
}



type Vertex struct {
	X,Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X *f
	v.Y = v.Y *f
}


func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f;
	v.Y = v.Y * f
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f <0 {
		return float64(-f)
	}
	return float64(f)
}

