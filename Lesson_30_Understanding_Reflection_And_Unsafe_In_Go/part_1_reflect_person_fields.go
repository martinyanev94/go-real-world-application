package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Alice", 30}
    t := reflect.TypeOf(p)

    fmt.Println("Type:", t.Name())
    fmt.Println("Number of fields:", t.NumField())
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field %d: %s (%s)\n", i, field.Name, field.Type)
    }
}
