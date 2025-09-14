package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    sharedVar := 1

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            sharedVar += i // This may cause a race condition
            fmt.Println("Goroutine:", i, "Shared Var:", sharedVar)
        }(i)
    }
    
    wg.Wait()
}
