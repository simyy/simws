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

    fmt.Println(tair.Count())

    fmt.Println(tair.Find([]byte{'a'}))

    fmt.Println("==> end")
}
