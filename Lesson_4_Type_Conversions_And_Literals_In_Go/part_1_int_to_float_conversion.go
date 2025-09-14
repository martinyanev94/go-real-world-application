package main

import "fmt"

func main() {
    var integerValue int = 42
    var floatValue float64 = float64(integerValue) // Explicitly converting int to float64
    fmt.Printf("Integer Value: %d, Float Value: %.2f\n", integerValue, floatValue)
}
