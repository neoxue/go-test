package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dict := map[string]int{"foo": 1, "bar": 2}
	value, ok := dict["baz"]
	if ok {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("key not found")
	}

	jsonstr := `
	{
		"a": {
			"b":"d"
		},
		"b":[
			"ddd",
			{
				"eee":"fff"
			}
		]
	}
	`

	fmt.Println(jsonstr)
	var a map[string]interface{}
	err := json.Unmarshal([]byte(jsonstr), &a)
	fmt.Println(err)
	aa := a["a"]
	fmt.Println(aa)
	aa.(map[string]interface{})["b"] = "bb"
	fmt.Println(aa)
	fmt.Println(a)
	testb(aa.(map[string]interface{}))
	fmt.Println(aa)
	fmt.Println(a)

	b := a["b"].([]interface{})
	b[0] = "ddd1"
	fmt.Println(a)

	c := b[1]
	c1 := c.(map[string]interface{})
	c1["ccc"] = "jjj"
	fmt.Println(a)

	jsonnew := `
	{
		"a":null
	}
	`
	var ax map[string]interface{}
	err1 := json.Unmarshal([]byte(jsonnew), &ax)
	fmt.Println(err1)
	fmt.Println(ax)
	fmt.Println(ax["a"] == nil)

}
func testb(amap map[string]interface{}) {
	amap["b"] = "dddd"

}
