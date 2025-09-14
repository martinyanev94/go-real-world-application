package main

import "fmt"

type operation func(int, int) int

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func operate(op operation, a, b int) int {
    return op(a, b)
}

func main() {
    fmt.Println("Add:", operate(add, 5, 3))          // Output: Add: 8
    fmt.Println("Subtract:", operate(subtract, 5, 3)) // Output: Subtract: 2
}
