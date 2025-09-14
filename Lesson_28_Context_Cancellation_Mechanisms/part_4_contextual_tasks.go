package main

import (
    "context"
    "fmt"
    "time"
)

func task1(ctx context.Context) {
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("Task 1 completed")
    case <-ctx.Done():
        fmt.Println("Task 1 cancelled:", ctx.Err())
    }
}

func task2(ctx context.Context) {
    select {
    case <-time.After(4 * time.Second):
        fmt.Println("Task 2 completed")
    case <-ctx.Done():
        fmt.Println("Task 2 cancelled:", ctx.Err())
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go task1(ctx)
    
    ctx2, cancel2 := context.WithTimeout(ctx, 2*time.Second)
    defer cancel2()

    go task2(ctx2)

    time.Sleep(3 * time.Second)
  
    cancel() // Canceling the parent context
    time.Sleep(1 * time.Second) // Give tasks time to terminate
}
