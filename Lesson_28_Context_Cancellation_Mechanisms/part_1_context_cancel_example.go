package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    
    go func() {
        time.Sleep(2 * time.Second)
        cancel() // Cancelling the context after 2 seconds
    }()

    select {
    case <-ctx.Done():
        fmt.Println("Operation cancelled:", ctx.Err())
    case <-time.After(5 * time.Second):
        fmt.Println("Operation completed successfully")
    }
}
