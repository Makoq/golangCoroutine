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
var ch chan int

func func1(ch chan int) {
	time.Sleep(4 * time.Second)
	ch <- 1
	defer wg.Done()
}

func func2(ch chan int) {
	time.Sleep(6 * time.Second)

	ch <- 2
	defer wg.Done()

}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg.Add(2)
	go func1(ch1)
	go func1(ch2)
	wg.Wait()

	//for+select 持续监听select
	//这里会监听到func1和func2都执行完成后结束
	for {
		select {
		case <-ch1:
			fmt.Println("ch1")
		case <-ch2:
			fmt.Println("ch2")

		}

	}
}
