package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	time.Sleep(1 * time.Second)
	x,y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x,y = y, y+x
	}
	close(c)
}

func main () {
	c := make (chan int, 10)
	fmt.Println(cap(c))
	go fibonacci(cap(c), c)
	fmt.Println(cap(c))
	for i := range c {
		fmt.Println(i)
	}
}