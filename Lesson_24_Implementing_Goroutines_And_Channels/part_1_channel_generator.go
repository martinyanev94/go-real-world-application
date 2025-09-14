package main

import (
    "fmt"
)

func generateNumbers(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i // Send value to channel
    }
    close(ch) // Close the channel when done
}

func main() {
    ch := make(chan int) // Create a channel

    go generateNumbers(ch) // Start the generator goroutine

    for num := range ch { // Receive data from the channel
        fmt.Println(num) // Print the received number
    }
}
