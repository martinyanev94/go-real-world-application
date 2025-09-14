package main

import (
    "encoding/json"
    "fmt"
    "log"
)

func main() {
    jsonData := []byte(`[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25, "city": "New York"}
    ]`)

    var people []map[string]interface{}
    err := json.Unmarshal(jsonData, &people)
    if err != nil {
        log.Fatalf("Error parsing JSON: %s", err)
    }

    for _, person := range people {
        fmt.Printf("Name: %s, Age: %v\n", person["name"], person["age"])
        if city, ok := person["city"]; ok {
            fmt.Printf("City: %s\n", city)
        } else {
            fmt.Println("City: Not specified")
        }
    }
}
