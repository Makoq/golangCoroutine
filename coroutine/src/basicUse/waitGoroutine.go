package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Hello() {
	defer wg.Done()
	time.Sleep(4 * time.Second)

	fmt.Println("hello everybody , I'm asong")
}

func main() {

	go Hello()
	wg.Add(1)

	fmt.Println("Golang梦工厂")
	wg.Wait()

}
