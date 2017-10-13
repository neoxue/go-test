package main

import "fmt"

func main () {
	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v %v %v %q %v %q\n", i, f, b, s, s, b)
	fmt.Printf("%v", s)
	fmt.Printf("%q", i)

}
