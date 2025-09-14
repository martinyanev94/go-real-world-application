package main

import (
    "fmt"
    "os"
)

type ConfigError struct {
    Filename string
    Err      error
}

func (e *ConfigError) Error() string {
    return fmt.Sprintf("config error with file %s: %v", e.Filename, e.Err)
}

func (e *ConfigError) Unwrap() error {
    return e.Err
}

func readConfig(filename string) error {
    _, err := os.ReadFile(filename)
    if err != nil {
        return &ConfigError{Filename: filename, Err: err} // Returning custom error
    }
    return nil
}

func handleError(err error) {
    if errors.As(err, &ConfigError{}) {
        fmt.Println("Custom error detected:", err)
    } else {
        fmt.Println("An error occurred:", err)
    }
}

func main() {
    err := readConfig("nonexistent.txt")
    handleError(err)
}
