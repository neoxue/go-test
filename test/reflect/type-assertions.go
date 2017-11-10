package main

import "fmt"

func main() {
	var i interface{} = "hello"
	af := i.(float64)
	fmt.Println(af)
	return
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)

}
