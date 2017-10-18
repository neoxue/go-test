package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	time.Sleep(100 * time.Millisecond)
	sum := 0;
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main () {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Println(len(s)/2)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <- c
	fmt.Println(x, y, x+y)
}
