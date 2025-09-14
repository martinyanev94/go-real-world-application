func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2 // Simulate work by doubling the job value
    }
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send 10 jobs
    for j := 1; j <= 10; j++ {
        jobs <- j
    }
    close(jobs) // Close jobs channel

    // Collect results
    for a := 1; a <= 10; a++ {
        fmt.Println(<-results) // Print results from workers
    }
}
