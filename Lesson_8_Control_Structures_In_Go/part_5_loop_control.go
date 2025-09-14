package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            fmt.Println("Skipping number 5")
            continue
        }
        if i == 8 {
            fmt.Println("Breaking at number 8")
            break
        }
        fmt.Println("Current number:", i)
    }
}
