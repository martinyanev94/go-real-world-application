package main

import (
	"fmt"
	"time"
)

// simulateWork simulates some heavy work by allocating memory
func simulateWork() {
	_ = make([]byte, 1<<20) // Allocating 1MB
	time.Sleep(10 * time.Millisecond)
}

func main() {
	for i := 0; i < 10; i++ {
		simulateWork() // This operation will increase GC activity
	}
}
