package main

import (
    "fmt"
    "unsafe"
)

type Data struct {
    Value  uint32
    Label  [10]byte
    Active bool
}

func main() {
    bytes := [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
    data := *(*Data)(unsafe.Pointer(&bytes))
    fmt.Println("Value:", data.Value)
    fmt.Println("Label:", string(data.Label[:]))
    fmt.Println("Active:", data.Active)
}
