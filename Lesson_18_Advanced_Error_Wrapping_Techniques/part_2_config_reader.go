package main

import (
    "errors"
    "fmt"
    "os"
)

func readConfig(filename string) error {
    _, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("failed to read config file %s: %w", filename, err)
    }
    return nil
}

func handleError(err error) {
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("The specified config file does not exist.")
    } else {
        fmt.Println("An error occurred:", err)
    }
}

func main() {
    err := readConfig("nonexistent.txt")
    handleError(err)
}
