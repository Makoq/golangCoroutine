package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Hello(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("done")
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	default:
		fmt.Println("default")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	go Hello(ctx)
	time.Sleep(4 * time.Second)
}
