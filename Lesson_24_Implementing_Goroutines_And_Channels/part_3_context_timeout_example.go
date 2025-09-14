package main

import (
    "context"
    "fmt"
    "time"
)

func doWork(ctx context.Context) {
    select {
    case <-time.After(3 * time.Second): // Simulate work with a timeout
        fmt.Println("Work completed")
    case <-ctx.Done(): // Respond to context cancellation
        fmt.Println("Work canceled")
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel() // Cancel the context when done

    go doWork(ctx) // Start the work in a goroutine

    time.Sleep(2 * time.Second) // Allow some time for the work to process
}
