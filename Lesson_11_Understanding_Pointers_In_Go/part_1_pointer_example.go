package main

import "fmt"

func updateValue(val *int) {
    *val = 100 // Dereferencing to change the value at the address of val
}

func main() {
    var a int = 50
    fmt.Println("Original value of a:", a)
    updateValue(&a) // Pass the address of a
    fmt.Println("Updated value of a:", a)
}
