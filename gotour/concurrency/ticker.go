package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan map[string]interface{}, 100)
	go funca(ch)
	go funcconsume(ch)
	time.Sleep(100 * time.Second)
}

func funcconsume(ch chan map[string]interface{}) {
	for {
		a := <-ch
		fmt.Println(a)
	}

}
func funca(ch chan map[string]interface{}) {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		ch <- map[string]interface{}{"a": t}
		fmt.Println("Tick at", t)
	}
}
