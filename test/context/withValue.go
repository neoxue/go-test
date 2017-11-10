package main

import (
	"context"
	"fmt"
	"time"
)

type favContextKey string

func f(ctx context.Context, k favContextKey) {
	if v := ctx.Value(k); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", k)
}

func main() {
	k := favContextKey("language")
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
	ctx2 := context.WithValue(ctx, k, "Go")
	ctx3, _ := context.WithTimeout(ctx2, 50*time.Millisecond)
	testvalue(ctx3, "language")
	testvalue(ctx3, "color")
	time.Sleep(2 * time.Second)
	testvalue(ctx3, "langguage")
	testvalue(ctx3, "language")
}

func testvalue(ctx context.Context, key string) {
	f(ctx, favContextKey(key))
}
