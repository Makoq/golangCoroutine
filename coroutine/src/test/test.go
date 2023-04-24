package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

}

//读取指定路径文件目录的测试方法
func ReadDirTest() {
	re, _ := ioutil.ReadDir("")
	for i, r := range re {
		fmt.Println(i, r.Name())
	}
}
