package main

import "fmt"

func main () {

	fmt.Println(c())
	b()
	a()
	fmt.Println("\n")
	testpanic()
}
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func testpanic() {

	defer fmt.Println("it's panic")
	a := 0;
	fmt.Println(1/a)
	return
}