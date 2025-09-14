package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("item not found")

func findItem(name string) error {
    if name != "valid" {
        return ErrNotFound
    }
    return nil
}

func main() {
    err := findItem("invalid")
    if errors.Is(err, ErrNotFound) {
        fmt.Println("Handle not found error")
    }
}
