package main

import "fmt"

func main() {
    arr := []int{10, 20, 30, 40, 50}
    for i := 0; i < len(arr); i++ {
        increment(&arr[i]) // Pass the address of each element
    }
    fmt.Println("Updated array:", arr)
}

func increment(val *int) {
    *val++ // Increment the value at the address
}
