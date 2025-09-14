package main

import (
    "fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2 // Simulate processing
    }
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    for w := 1; w <= 3; w++ { // Start three workers
        go worker(w, jobs, results)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j // Send jobs to workers
    }
    close(jobs) // Close jobs channel

    for a := 1; a <= 5; a++ {
        fmt.Println(<-results) // Receive results from workers
    }
}
