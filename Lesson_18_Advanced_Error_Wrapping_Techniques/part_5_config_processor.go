package main

import (
    "errors"
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
        return &ConfigError{Filename: filename, Err: err}
    }
    return nil
}

func validateUserInput(input string) error {
    if len(input) == 0 {
        return errors.New("input cannot be empty")
    }
    return nil
}

func processConfig(filename, userInput string) error {
    if err := readConfig(filename); err != nil {
        return fmt.Errorf("failed in processConfig: %w", err)
    }
    
    if err := validateUserInput(userInput); err != nil {
        return fmt.Errorf("user input validation failed: %w", err)
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
    err := processConfig("nonexistent.txt", "")
    handleError(err)
}
