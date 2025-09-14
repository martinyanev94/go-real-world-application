package main

import "fmt"

func main() {
    var a int = 42
    var p *int = &a // p holds the memory address of a

    fmt.Println("Value of a:", a)
    fmt.Println("Address of a:", &a)
    fmt.Println("Value of p (address of a):", p)
    fmt.Println("Value stored at address p:", *p) // Dereferencing p to get the value of a
}
