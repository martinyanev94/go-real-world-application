package main

import "fmt"

func calculate(a int, b int) (int, int) {
    sum := a + b
    difference := a - b
    return sum, difference
}

func main() {
    sum, difference := calculate(10, 4)
    fmt.Println("Sum:", sum, "Difference:", difference)
}
