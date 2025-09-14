package main

import "fmt"

func riskyFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    panic("Oh no! Something went wrong...")
}

func main() {
    fmt.Println("Before risky function call")
    riskyFunction()
    fmt.Println("After risky function call")
}
