package main

import (
	"fmt"
	"sync"
)

/*
单向通道：
	只发送通道，只写，out chan<- int
	只接收通道，只读，in <-chan int
*/
var wg sync.WaitGroup

func forIn(out chan<- int) {
	out <- 1
	close(out)

}
func forOut(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i
	}
	close(out)

}
func print(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
func main() {

	// 只发送通道,只写
	w := make(chan int)
	// 只接收通道，只读
	r := make(chan int)

	go forIn(w)
	go forOut(r, w)

	print(w)

}
