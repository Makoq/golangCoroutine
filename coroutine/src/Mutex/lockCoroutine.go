package main

import (
	"fmt"
	"sync"
)

/**
*	互斥锁用法:访问共享资源时避免冲突，等待其他协程访问后再访问
*
**/

var x int = 1
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	lock.Lock()
	x = x + 1
	fmt.Println(x)
	lock.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println("finally", x)
}
