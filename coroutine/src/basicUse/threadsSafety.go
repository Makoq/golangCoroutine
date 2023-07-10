package main

import (
	"fmt"
	"sync"
)

//多个线程访问同一个内存变量资源
//通过加锁来控制线程安全


type Counter struct{
	mu sync.Mutex
	count int
}

func(c *Counter) Increment(){
	c.mu.Lock()
	c.count=c.count+1
	defer c.mu.Unlock()
}

func(c *Counter) Count() int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main(){

	ct:=Counter{}
	var wg sync.WaitGroup

	ct.count=2

	for i:=0;i<10;i++{
		wg.Add(1)
		go func ()  {
			ct.Increment()
			wg.Done()
		}()
	}
	fmt.Println("count is : ",ct.Count())
}