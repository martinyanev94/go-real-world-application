func doWork(ctx context.Context) {
    select {
    case <-time.After(2 * time.Second): // Simulating work
        fmt.Println("Work completed")
    case <-ctx.Done(): // Handle context cancelation
        fmt.Println("Work canceled")
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    go doWork(ctx) // Start work in a goroutine

    time.Sleep(3 * time.Second) // Wait for a while
}
