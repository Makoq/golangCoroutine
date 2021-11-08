package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器
	tick := time.Tick(time.Second)
	for i := range tick {
		fmt.Println("sdsd", i)
	}
}
