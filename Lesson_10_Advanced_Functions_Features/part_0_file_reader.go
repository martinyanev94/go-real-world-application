package main

import (
    "fmt"
    "os"
)

func readFile(filePath string) {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() // Ensure that the file is closed after function returns

    var content = make([]byte, 100)
    _, err = file.Read(content)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println("File content:", string(content))
}

func main() {
    readFile("example.txt") // Make sure to replace with a valid file path
}
