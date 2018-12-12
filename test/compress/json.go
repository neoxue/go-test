package main

import (
	"encoding/json"
	"fmt"
	"github.com/neoxue/goutils"
	"reflect"
	"strconv"
)

func main() {
	str := "{\"a\":[1,2,3]}"
	b := map[string]interface{}{}
	json.Unmarshal([]byte(str), &b)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b["a"]))

	fmt.Println(goutils.InArray(1, b["a"].([]interface{})))

	fmt.Println(float64(1.0) == float64(1))

	fmt.Println(strconv.FormatFloat(1.01, 'f', -1, 64))
}
