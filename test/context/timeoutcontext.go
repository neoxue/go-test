package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func test2(is_first *bool, prev_map *map[string]string) {
	rand.Seed(int64(time.Now().Nanosecond()))
	ch := make(chan bool)
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)

	//defer cancel()

	go func() {
		if *is_first {
			(*prev_map)["a"] = "1"
		} else {
			(*prev_map)["a"] = "2"
		}
		fmt.Println(1)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println(2)
		ch <- true
	}()

	for {
		select {
		case a := <-ch:
			fmt.Println(a)
			*is_first = false
			(*prev_map)["b"] = "fast"
			return
		case a := <-ctx.Done():
			fmt.Println(a)
			*is_first = true
			(*prev_map)["b"] = "slow"
			return
		}
	}
}

func main() {
	is_first := true
	prev_map := make(map[string]string)
	for {
		test2(&is_first, &prev_map)
		time.Sleep(time.Second)
		log.Println(prev_map, is_first)
		fmt.Println("\n\n\n")
	}
}
