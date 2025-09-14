package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    mu sync.Mutex // Mutex to protect counter access
    count int
}

func (c *Counter) Increment() {
    c.mu.Lock() // Acquire the lock
    c.count++
    c.mu.Unlock() // Release the lock
}

func main() {
    counter := Counter{}
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ { // Start 100 goroutines
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("Final Counter:", counter.count) // Should print 100
}
