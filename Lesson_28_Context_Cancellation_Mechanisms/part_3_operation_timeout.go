package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel() // Ensure resources are cleaned up

    select {
    case <-time.After(5 * time.Second): // Simulated operation taking too long
        fmt.Println("Operation completed successfully")
    case <-ctx.Done():
        fmt.Println("Operation cancelled due to timeout:", ctx.Err())
    }
}
