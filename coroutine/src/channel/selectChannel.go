package main

import (
	"fmt"
	"sync"
	"time"
)

/**
*	无缓冲通道: ch := make(chan int)
*	有缓冲通道: ch := make(chan int, n) n表示通信的元素个数
**/
var wg sync.WaitGroup

func func1(ch chan int) {
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()
	defer wg.Done()
}

func func2(ch chan int) {
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 2
	}()
	defer wg.Done()
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)
	go func1(ch1)
	go func2(ch2)
	wg.Wait()
	timeout := time.After(4 * time.Second)

	var msg1 int
	var msg2 int
	var received bool

loop:
	for {
		select {
		case msg1 = <-ch1:
			received = true
		case msg2 = <-ch2:
			received = true
		case <-timeout:
			fmt.Println("timeout")
			break loop
		default:
			fmt.Println("no message")
			time.Sleep(100 * time.Millisecond)
		}

		if received {
			fmt.Println("xxxxxxxxx", msg1)
			fmt.Println("xxxxxxxxx", msg2)
			break loop
		}
	}
}
