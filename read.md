知识点学习
- 基础知识
  - Func
    - defer
      - return前执行
      - 多个类似于栈，后面的先执行 详情
    - 多个返回值
```   
func exits(m map[string]string,k string)(v string,ok string){
    v,ok=m[k]
    return v,ok
}
```
- 结构体方法
```
type User struct{
    name string
    password number
}
func (u User) checkUsr(u User,pwd string) bool{
    return u.password==pwd
}
```

- Slice
    - slice、array区别
    - golang里是值传递
    - slice扩容逻辑和坑
    - append时，会根据cap扩容时决定是否重新申请内存，所以最好在append时使用三个参数，以保证原始slice不变，详情
      - 切片的结构体定义是长度、容量和一个指向数组的指针。
      - 当加入元素超出容量时会发生扩容，扩容会导致重新申请内存形成新的切片，此时新的切片和扩容前的切片互相独立；append(aee[:2:2],el...)的用法就是让扩容发生，产生新的切片，独立于原始切片
```
func _fn(nums []int, sum int, tp []int, all *[][]int) {
    if len(tp) >= 3 && sum != 0 {
        return
    } else if len(tp) == 3 && sum == 0 {
        sort.Ints(tp)
        if arrHas(*all, tp) != true {
            *all = append(*all, tp)
        }
        return
    } else {
        for i := 0; i < len(nums); i++ {
            next_sum := append(nums[:i:i], nums[i+1:]...)
            _sum := sum + nums[i]
            _tp := append(tp, nums[i])
            _fn(next_sum, _sum, _tp, all)
        }
    }
}
```
- copy(target,source)
      - 初始化target长度
      - 二维数组拷贝时，其实每个元素都是引用所以要对子元素进行拷贝
```
func _copy(a [][]byte,b [][]byte)[][]byte{
    for i,_:=range b{
        arr:=make([]byte,len(b[i]))
        copy(arr,b[i]) 
        a[i]=arr
    }
    return a
}
```
  - map
    - map传值是值传递，传的是指针
    - map修改会影响到原始的值
    - map不可直接去值（&map）,map会随着扩容而改变内存地址
    - map存的是值，取值时发生的是copy
    - map删除key，不会自动缩容 
  - channel
    - channel的使用，必须是需要多个goroutine之间传递数据的场景，如果是单纯的队列，用slice就好了
    - channel有锁
    - channel底层是个ringBuffer
    - channel调用会触发调度
    - channel不适用于高并发、高性能编程
    - for+select会造成死循环
    - select 的break不会跳出for循环
    - 是否定义长度
      - 有长度：有缓冲区通道（buffered channel）
        - 会发生两次copy
          - sendG->buffer
          - buffer->receiveG
      - 没长度：无缓冲区通道（unbuffered channel），必须有接受协程
        - 只有一次copy
          - sendG->bufferG
        - Unbuffered channel receive返回后，send才返回，这就造成，无缓冲通道必须要有接受的协程，否则报错
  - Struct
    - struct{} 空结构不占内存
  - 错误处理
```
import "errors"
func exists(m map[string]string,pwd string) (re string,err error){
    password,ok=m[pwd]
    if ok==false{
        return nil, errors.new("no password")
    }
    return re,nil
}
```
- 进阶知识
  - sync包
    - sync.WaitGroup
等待一组同步协程的执行结束。
    - 主要有三个方法
      - sync.WaitGroup.wait: 等待协程执行结束
      - sync.WaitGroup.Add(delta int):  添加协程数目
      - sync.WaitGroup.Done: 声明协程执行结束
    - 并发请求接口用法

```
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "sync"
)

/**
*   协程基本用法
*   sync.WaitGroup.wait: 等待协程执行结束
*   sync.WaitGroup.Add:  添加协程数目
*   sync.WaitGroup.Done: 声明协程执行结束
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
```
- sync.Mutex. 互斥锁
使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
 - 基本用法
mutex := &sync.Mutex{}

mutex.Lock()
// Update共享变量 (比如切片，结构体指针等)
// 比如下例子中的x
mutex.Unlock()

- 用法示例
```
package main


import (
    "fmt"
    "sync"
)


/**
*   两个协程访问同一个内存块，就会出现竞态问题
*   互斥锁用法:访问共享资源时避免冲突，等待其他协程访问后再访问
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
```
- sync.RWMutex. 读写互斥锁
一般用在大量读操作、少量写操作的情况。
同一时刻可以有多个读，但只能有一个写

Context 包
提供一个请求从API请求边界到各goroutine的请求域数据传递、取消信号及截止时间等能力
- Context.Background
- Context.WithDeadline
- Context.WithTimeout
- Context.WithCancel

TIme包
- Tick
- After
- Sleep

Os包
- File
  - Create
    - 创建
ioutil
- readAll
  - 读取全部内容，比如读取response.body

IO包
- Writer和MutilWriter：写入一个和多个数据流中
  - 在 Go 语言中，io.Writer 和 io.MultiWriter 都是接口类型，用于写入数据到某个数据流中。它们的使用场景有一些不同。
  - io.Writer 接口包含一个 Write 方法，用于将字节数据写入到某个数据流中。通常情况下，我们可以将一个 io.Writer 实现作为参数传递给一个函数，以便在函数中将数据写入到该数据流中。例如：
func writeTo(w io.Writer) error {
    // 将数据写入到 w 中
    _, err := w.Write([]byte("hello world"))
    return err
}

  - io.MultiWriter 接口同样包含一个 Write 方法，用于将字节数据写入到多个数据流中。它可以接收多个 io.Writer 实现作为参数，并将数据写入到所有这些实现中。例如：
func writeToMultiple(w1 io.Writer, w2 io.Writer) error {
    // 将数据同时写入到 w1 和 w2 中
    mw := io.MultiWriter(w1, w2)
    _, err := mw.Write([]byte("hello world"))
    return err
}

  - 可以看到，io.MultiWriter 的主要作用是将数据写入到多个数据流中，而不是单个数据流。它在需要同时向多个数据流中写入数据的情况下非常有用，例如将数据同时写入到文件和网络连接中。

关于测试

  - 回归测试
    -   qa同学手动回归功能
  - 集成测试
    - 系统功能层面的测试
  - 单元测试
    - 接口、函数的测试
go语言层面的性能优化
  - Benchmark
    - 以Benchmark开头的函数
    - go test -bench . 执行当前模块下的测试函数
    - 用法例子
```
// fib_test.go
package main

import "testing"

func BenchmarkFib(b *testing.B) {
        for n := 0; n < b.N; n++ {
                fib(30) // run fib(30) b.N times
        }
}
```
- Slice
    - 预声明len
    - 在一个原较大切片上截取一个较小的子切片，优先用copy，而不是sub:=par[a:b]
  - 

