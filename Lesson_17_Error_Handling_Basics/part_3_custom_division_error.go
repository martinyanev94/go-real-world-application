package main

import (
    "fmt"
)

// Custom error type
type DivisionError struct {
    numerator   int
    denominator int
}

// Implementing the error interface
func (e DivisionError) Error() string {
    return fmt.Sprintf("cannot divide %d by %d", e.numerator, e.denominator)
}

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, DivisionError{numerator, denominator} // Returning custom error
    }
    return numerator / denominator, numerator % denominator, nil
}

func main() {
    numerator := 20
    denominator := 0
    _, _, err := calcRemainderAndMod(numerator, denominator)
    if err != nil {
        fmt.Println(err) // Output: cannot divide 20 by 0
    }
}
