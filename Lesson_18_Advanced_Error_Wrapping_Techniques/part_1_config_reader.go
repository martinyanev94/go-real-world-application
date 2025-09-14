package main

import (
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

func main() {
    err := readConfig("nonexistent.txt")
    if err != nil {
        fmt.Println(err) // Output: failed to read config file nonexistent.txt: open nonexistent.txt: no such file or directory
    }
}
