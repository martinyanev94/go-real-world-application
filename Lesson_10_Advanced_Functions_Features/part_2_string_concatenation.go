package main

import (
    "fmt"
    "strings"
)

func concatenate(separator string, words ...string) string {
    return strings.Join(words, separator)
}

func main() {
    result := concatenate(", ", "Hello", "world", "from", "Go")
    fmt.Println(result)
}
