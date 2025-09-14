package main

import (
    "encoding/json"
    "errors"
    "fmt"
)

type Employee struct {
    Name     string `json:"name"`
    Position string `json:"position"`
}

func (e *Employee) UnmarshalJSON(data []byte) error {
    type Alias Employee
    alias := &Alias{}
    if err := json.Unmarshal(data, alias); err != nil {
        return err
    }
    
    if alias.Name == "" || alias.Position == "" {
        return errors.New("name and position must be provided")
    }
    
    e.Name = alias.Name
    e.Position = alias.Position
    return nil
}

func main() {
    jsonData := []byte(`{"name": "Alice", "position": ""}`)

    var employee Employee
    err := json.Unmarshal(jsonData, &employee)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Printf("%s is a %s.\n", employee.Name, employee.Position)
}
