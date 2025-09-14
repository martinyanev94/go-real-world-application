package main

import "fmt"

func main() {
    var defaultInt int      // zero value for int is 0
    var defaultFloat float64 // zero value for float64 is 0.0
    var defaultString string // zero value for string is ""

    fmt.Printf("Default Int: %d\n", defaultInt)
    fmt.Printf("Default Float: %.2f\n", defaultFloat)
    fmt.Printf("Default String: %q\n", defaultString)
}
