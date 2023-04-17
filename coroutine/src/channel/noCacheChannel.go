package main

import (
	"fmt"
)

/**
*	无缓冲通道: ch := make(chan int)
*	有缓冲通道: ch := make(chan int, n) n表示通信的元素个数
**/

var ch chan int

func recv(ch chan int) {
	re := <-ch
	fmt.Println("receive success", re)
	close(ch)
}

func main() {
	ch := make(chan int)
	//无缓冲通道，必须要建立一个协程去接收通道的数据
	go recv(ch)
	ch <- 10
	fmt.Println("send success")
}
