package main

/*
import (
	"sync"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"net/http"
)

func main()  {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()

		ss := make(os.Signal, 1)
		signal.Notify(ss, syscall.SIGINT, syscall.SIGTERM)
		for s := ss {
			fmt.Println("Got signal", s)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		svr := &http.Server{Addr:":8080", Handler:nil}
		fmt.Println(svr.ListenAndServe())

	}()

}
*/