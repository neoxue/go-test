package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	//teststop()

	fmt.Println("Wait signal:") //Got signal: terminated
	signal.Notify(c)
	fmt.Println("Waiting signal:") //Got signal: terminated
	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s) //Got signal: terminated
}

func teststop() {
	c := make(chan os.Signal, 0)
	signal.Notify(c)

	signal.Stop(c) //不允许继续往c中存入内容
	s := <-c       //c无内容，此处阻塞，所以不会执行下面的语句，也就没有输出
	fmt.Println("Got signal:", s)
}
