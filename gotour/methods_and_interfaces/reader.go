package main

import (
	"strings"
	"fmt"
	"io"
)

func main () {
	r := strings.NewReader("hello, world!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b= %v", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		break;

		if err == io.EOF {
			break;
		}
	}
}
