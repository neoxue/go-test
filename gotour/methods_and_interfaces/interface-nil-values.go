package main

import (
	"fmt"
)

type I4 interface {
	M4()
}
type T44 struct {
	S string
}


func main() {
	var i I4
	describe4(i)
	i.M4()
}

func (t *T44)M4 () {
	if (t == nil) {
		fmt.Println("<nil>")
	}
	return;
}

func describe4(i I4) {
	fmt.Printf("$v, %T", i, i)
}
