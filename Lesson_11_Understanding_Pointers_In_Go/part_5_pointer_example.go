package main

import "fmt"

func main() {
    a := 42
    p := &a
    fmt.Printf("Value of a: %d, Address of a: %p, Value of p (address of a): %p\n", a, &a, p)
}
