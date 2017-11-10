package main

import "fmt"

type a interface {
}

type b interface {
}

func main() {
	var testa *b
	_, testa = testb()
	if testa != nil {
		fmt.Println("yes")
	}

}

func testb() (string, *b) {
	return "", nil
}
