package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func add(url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	//有效返回时进行处理
	if resp != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
		resp.Body.Close()
	}

}

func main() {
	url := "http://localhost:8082/common/user"
	//一百万次并发请求
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go add(url)
	}
	wg.Wait()
}
