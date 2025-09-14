package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go printNumbers() // Launching the goroutine

    // Keep the main function alive to see the output
    time.Sleep(6 * time.Second)
}
