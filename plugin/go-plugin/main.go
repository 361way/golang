package main

import (
        "fmt"
        "plugin"
)

func main() {
        p, _ := plugin.Open("./aplugin.so")
        add, _ := p.Lookup("Add")
        sub, _ := p.Lookup("Subtract")
        sum := add.(func(int, int) int)(11, 2)
        fmt.Println(sum)
        subt := sub.(func(int, int) int)(11, 2)
        fmt.Println(subt)
}
