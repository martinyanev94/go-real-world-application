package main

import (
    "fmt"
    "strings"
)

func main() {
    original := "Hello, Gopher!"
    uppercased := strings.ToUpper(original)
    replaced := strings.ReplaceAll(original, "Gopher", "World")

    fmt.Println("Original:", original)
    fmt.Println("Uppercased:", uppercased)
    fmt.Println("Replaced:", replaced)
}
