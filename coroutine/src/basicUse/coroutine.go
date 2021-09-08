package main

import (
	"fmt"
	"sync"
)

/**
*	协程基本用法
*	sync.WaitGroup.wait: 等待协程执行结束
*	sync.WaitGroup.Add:  添加协程数目
*	sync.WaitGroup.Done: 声明协程执行结束
*
**/

var wg sync.WaitGroup

func add(x, y int) {
	defer wg.Done()
	z := x + y
	fmt.Println(z)

}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(i, i)
	}
	wg.Wait()
}
