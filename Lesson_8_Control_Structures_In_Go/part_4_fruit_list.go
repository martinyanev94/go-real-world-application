package main

import "fmt"

func main() {
    fruits := []string{"apple", "banana", "cherry"}

    for index, fruit := range fruits {
        fmt.Printf("Fruit at index %d is %s\n", index, fruit)
    }
}
