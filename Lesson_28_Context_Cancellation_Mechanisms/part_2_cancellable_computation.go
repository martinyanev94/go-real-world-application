package main

import (
    "context"
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func longComputation(ctx context.Context, input int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < input; i++ {
        select {
        case <-ctx.Done():
            fmt.Println("Computation cancelled")
            return
        default:
            // Simulating a time-consuming task
            time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
            fmt.Printf("Calculating... %d\n", i)
        }
    }
    fmt.Println("Computation completed")
}

func main() {
    var wg sync.WaitGroup
    
    ctx, cancel := context.WithCancel(context.Background())
    
    input := 10
    wg.Add(1)
    go longComputation(ctx, input, &wg)

    time.Sleep(500 * time.Millisecond) // Wait for a bit before canceling
    
    cancel() // Cancelling the context
    wg.Wait() // Waiting for the computation to finish
}
