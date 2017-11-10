package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`[^\d]`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	url := "/book/list/5"
	d := re.ReplaceAllString(url, "")
	fmt.Println(d)
}
