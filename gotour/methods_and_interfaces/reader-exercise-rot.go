package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func main () {
	s := strings.NewReader("Lbh ssssss aaa sss!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
