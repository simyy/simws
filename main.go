package main

import (
    "fmt"
    "github.com/yxd123/simWS/tair"
)

func main() {
    fmt.Println("==> start")

    tair := tair.New()

    fmt.Println(tair.Count())

    tair.Add([]byte{'a'})
    tair.Add([]byte{'b'})
    tair.Add([]byte{'a', 'b'})

    fmt.Println(tair.Count())

    fmt.Println(tair.Find([]byte{'a'}))
    fmt.Println(tair.Find([]byte{'b'}))
    fmt.Println(tair.Find([]byte{'a', 'b'}))
    fmt.Println(tair.Find([]byte{'c'}))
    fmt.Println(tair.Find([]byte{'a', 'c'}))

    fmt.Println("==> end")
}
