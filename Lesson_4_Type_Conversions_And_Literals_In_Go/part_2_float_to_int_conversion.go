package main

import "fmt"

func main() {
    var floatValue float64 = 67.89
    var integerValue int = int(floatValue) // Converting float64 to int, which truncates the decimal
    fmt.Printf("Float Value: %.2f, Converted Integer Value: %d\n", floatValue, integerValue)
}
