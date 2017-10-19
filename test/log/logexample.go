package main

import (
	"github.com/pkg/errors"
	"log"
)

func main() {
	err := errors.New("test")
	log.Fatal(err)
}
