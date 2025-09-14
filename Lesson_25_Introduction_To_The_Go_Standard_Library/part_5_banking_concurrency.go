package main

import (
    "fmt"
    "sync"
)

var (
    balance int
    mu      sync.Mutex
)

func deposit(amount int, wg *sync.WaitGroup) {
    mu.Lock()
    balance += amount
    mu.Unlock()
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go deposit(100, &wg)
    go deposit(200, &wg)

    wg.Wait()
    fmt.Println("Balance:", balance)
}
