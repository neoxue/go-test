package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}
	s := S{}
	st := reflect.TypeOf(s)
	fmt.Println(st)
	field := st.Field(0)
	fmt.Println(field.Tag)
	fmt.Println(reflect.TypeOf(field.Tag))
	fmt.Println(reflect.TypeOf(field.Tag))
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
	fmt.Println(field.Tag.Get("color") == "blue")

}
