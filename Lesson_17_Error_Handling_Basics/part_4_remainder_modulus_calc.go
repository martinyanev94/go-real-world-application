package main

import (
    "errors"
    "fmt"
)

// Function to calculate the remainder and modulus
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, errors.New("denominator is zero")
    }
    return numerator / denominator, numerator % denominator, nil
}

// Function that wraps an error
func doCalculation(num, denom int) error {
    _, _, err := calcRemainderAndMod(num, denom)
    if err != nil {
        return fmt.Errorf("calculation error: %w", err)
    }
    return nil
}

func main() {
    err := doCalculation(20, 0)
    if err != nil {
        fmt.Println(err) // Output: calculation error: denominator is zero
        if errors.Is(err, errors.New("denominator is zero")) {
            fmt.Println("Specific error: Denominator cannot be zero!")
        }
    }
}
