package main

import (
    "fmt"
    "github.com/yxd123/simWS/tire"
)

func main() {
    fmt.Println("==> start")

    tair := tair.New()

    fmt.Println(tair.Count())

    //tire.Add([]byte{'b'})
    //tire.Add([]byte{'a', 'b'})

    //fmt.Println(tair.Count())

    //fmt.Println(tire.Find([]byte{'a'}))
    //fmt.Println(tire.Find([]byte{'b'}))
    //fmt.Println(tire.Find([]byte{'a', 'b'}))

    tire.Add([]byte("测试"))
    tire.Add([]byte("测试1"))
    tire.Add([]byte("测试2"))

    fmt.Println(tire.Count())

    fmt.Println(tire.Find([]byte("测")))
    fmt.Println(tire.Find([]byte("测试")))
    fmt.Println(tire.Find([]byte("测试1")))
    fmt.Println(tire.Find([]byte("测试3")))

    fmt.Println("==> end")
}
