package main

import (
"fmt"
"reflect"
)

func AnythingToSlice(a interface{}) interface{} {
	fmt.Println(reflect.TypeOf(a))
	v := reflect.ValueOf(a)
	defer func() {
		fmt.Println("Recovering")
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		result := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = v.Index(i).Interface()
		}
		return result
	default:
		panic("not supported")
	}
}

func main() {
	fmt.Println(AnythingToSlice([...]int{1, 2, 3})) // array
	fmt.Println(AnythingToSlice([]int{1, 2, 3}))   // slice
	fmt.Println(AnythingToSlice("a"))   // slice
	fmt.Println(AnythingToSlice([]int{1, 2, 3}))   // slice

}

