package main

import (
    "errors"
    "fmt"
)

// Function to calculate the remainder and modulus of two integers
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, errors.New("denominator is 0")
    }
    return numerator / denominator, numerator % denominator, nil
}

func main() {
    numerator := 20
    denominator := 0
    remainder, mod, err := calcRemainderAndMod(numerator, denominator)
    if err != nil {
        fmt.Println(err) // Output: denominator is 0
        return
    }
    fmt.Printf("Remainder: %d, Modulus: %d\n", remainder, mod)
}
