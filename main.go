package main

import (
    "fmt"
    "github.com/yxd123/simWS/tair"
)

func main() {
    fmt.Println("==> start")

    tair := tair.New()

    fmt.Println(tair.Count())

    //tair.Add([]byte{'b'})
    //tair.Add([]byte{'a', 'b'})

    //fmt.Println(tair.Count())

    //fmt.Println(tair.Find([]byte{'a'}))
    //fmt.Println(tair.Find([]byte{'b'}))
    //fmt.Println(tair.Find([]byte{'a', 'b'}))

    tair.Add([]byte("测试"))
    tair.Add([]byte("测试1"))
    tair.Add([]byte("测试2"))

    fmt.Println(tair.Count())

    fmt.Println(tair.Find([]byte("测")))
    fmt.Println(tair.Find([]byte("测试")))
    fmt.Println(tair.Find([]byte("测试1")))
    fmt.Println(tair.Find([]byte("测试3")))

    fmt.Println("==> end")
}
