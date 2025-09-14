ch := make(chan int) // Create a new channel for int type
func sendData(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i // Send value to channel
    }
    close(ch) // Close the channel when done
}

func main() {
    ch := make(chan int)

    go sendData(ch) // Launch the sendData goroutine

    // Receive data from channel
    for num := range ch {
        fmt.Println(num) // Will print numbers sent from sendData
    }
}
