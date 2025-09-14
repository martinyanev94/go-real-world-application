package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    fmt.Printf("Hello %s!\n", name)
}

func main() {
    go sayHello("Alice")
    go sayHello("Bob")
    time.Sleep(1 * time.Second) // Wait for goroutines to finish
}
