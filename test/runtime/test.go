package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println(runtime.Breakpoint)

}
