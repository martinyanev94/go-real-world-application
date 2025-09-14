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

    // Create 100 goroutines to increment the counter
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("Final Counter:", counter.count)
}
