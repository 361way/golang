package main

import (
    "fmt"
    "runtime"
)

func init() {
    runtime.GOMAXPROCS(1)  //使用单核
    //runtime.GOMAXPROCS(4)  //使用多核,多核测试时要保证主机核数多于当前的设置数
}

func main() {
    exit := make(chan int)
    go func() {
        defer close(exit)
        go func() {
            fmt.Println("b")
        }()
    }()

    for i := 0; i < 4; i++ {
        fmt.Println("a:", i)

        if i == 1 {
            runtime.Gosched()  //切换任务
        }
    }
    <-exit

}
