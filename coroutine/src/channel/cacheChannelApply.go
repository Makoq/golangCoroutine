package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
// 起三个并发协程，最快的响应，塞入通道，
func main() {

	ch1 := make(chan string, 3)

	go func() {
		rs, GetErr := http.Get("https://catfact.ninja/breeds?limit=1")
		if GetErr != nil {
			fmt.Println(GetErr)
			return
		}
		defer rs.Body.Close()

		body, _ := ioutil.ReadAll(rs.Body)

		ch1 <- string(body)
	}()
	go func() {
		rs, GetErr := http.Get("http://localhost:8082/common/user")
		if GetErr != nil {
			fmt.Println(GetErr)
			return
		}
		defer rs.Body.Close()

		body, _ := ioutil.ReadAll(rs.Body)

		ch1 <- string(body)
	}()
	go func() {
		rs, GetErr := http.Get("https://catfact.ninja/fact")
		if GetErr != nil {
			fmt.Println(GetErr)
			return
		}
		defer rs.Body.Close()
		body, _ := ioutil.ReadAll(rs.Body)

		ch1 <- string(body)
	}()
	fmt.Println("0--", len(ch1))

	x := <-ch1

	fmt.Println("01--", len(ch1))

	fmt.Println("02--", x)
}
