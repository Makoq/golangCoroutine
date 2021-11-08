package main

import (
	"fmt"
	"sync"
)

/**
*   两个协程访问同一个内存块，就会出现竞态问题
*	互斥锁用法:访问共享资源时避免冲突，等待其他协程访问后再访问
*   加锁保证了一块内存再某时刻只会被一个协程访问到
**/

var x int = 1
var wg sync.WaitGroup
var lock sync.Mutex

func add() {

	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	fmt.Println("coroutine execute: ", x)

	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println("finally", x)
}
//加互斥锁后正常的输出为10001
//不加互斥锁后输出为5585（某一次）