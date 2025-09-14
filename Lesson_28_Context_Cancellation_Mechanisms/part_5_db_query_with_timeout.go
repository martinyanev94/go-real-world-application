package main

import (
    "context"
    "fmt"
    "time"
)

// Simulated DB call
func dbQuery(ctx context.Context, query string) {
    select {
    case <-time.After(5 * time.Second):
        fmt.Println("DB query completed for:", query)
    case <-ctx.Done():
        fmt.Println("DB query cancelled:", ctx.Err())
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    dbQuery(ctx, "SELECT * FROM users")
}
