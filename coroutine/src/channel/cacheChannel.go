package main

var ch chan int

func main() {
	//只要声明通道大小，就证明是有缓冲通道，表示通道可塞入的元素数目
	v:=make(chan int,1)
	v<- 10
	print("send ok")
}
