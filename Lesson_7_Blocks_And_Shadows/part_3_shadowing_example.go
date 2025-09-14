package main

import "fmt"

func main() {
    outerX := 20
    fmt.Println("Outer X:", outerX)

    innerX := add(outerX)
    fmt.Println("Value of Inner X returned from add function:", innerX)
}

func add(value int) int {
    // Directly using the argument value instead of shadowing
    return value + 10
}
