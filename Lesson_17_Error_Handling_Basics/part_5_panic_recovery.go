package main

import (
    "fmt"
)

// A function that simulates a panic situation
func riskyFunction() {
    panic("something went terribly wrong!") // This will trigger a panic
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r) // Handles the panic
        }
    }()
    riskyFunction()
    fmt.Println("After riskyFunction") // This line won't execute
}
