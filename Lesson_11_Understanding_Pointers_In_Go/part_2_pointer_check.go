package main

import "fmt"

func checkPointer(p *int) {
    if p == nil {
        fmt.Println("Pointer is nil, cannot dereference")
    } else {
        fmt.Println("Value at pointer:", *p)
    }
}

func main() {
    var p *int // p is a nil pointer
    checkPointer(p) // Check for nil
    var a int = 75
    p = &a
    checkPointer(p) // Now p is pointing to a
}
