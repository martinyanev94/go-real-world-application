package main

import (
    "fmt"
    "errors"
)

func divide(a int, b int) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return float64(a) / float64(b), nil
}

func main() {
    result, err := divide(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }

    result, err = divide(4, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
