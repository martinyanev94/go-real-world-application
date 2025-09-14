package main

import (
    "fmt"
    "unsafe"
)

type A struct {
    a bool
    b int64
}

type B struct {
    b int64
    a bool
}

func main() {
    fmt.Println("Size of A:", unsafe.Sizeof(A{})) // Output: Size of A: 16
    fmt.Println("Size of B:", unsafe.Sizeof(B{})) // Output: Size of B: 16
}
