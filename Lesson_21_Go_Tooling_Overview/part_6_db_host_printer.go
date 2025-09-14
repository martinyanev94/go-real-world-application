package main

import (
    "fmt"
    "os"
)

func main() {
    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "localhost"
    }
    fmt.Printf("Database host: %s\n", dbHost)
}
