package main

import (
    "fmt"
    "github.com/yxd123/simWS/tair"
)

func main() {
    fmt.Println("==> start")

    tair := tair.New()

    fmt.Println(tair.Count())

    fmt.Println("==> end")
}
