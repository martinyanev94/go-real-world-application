package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    jsonData, err := json.Marshal(person)
    if err != nil {
        log.Fatalf("Error occurred while marshaling: %s", err)
    }
    fmt.Println(string(jsonData))
}
