package main
import (
    "fmt"
)

func main() {

    go func() { //子协程   //没来的及执行主进程结束
        for i := 0; i < 5; i++ {
            fmt.Println("go")
        }
    }()

    for i := 0; i < 2; i++ { //默认先执行主进程主进程执行完毕
        fmt.Println("hello")
    }
}
