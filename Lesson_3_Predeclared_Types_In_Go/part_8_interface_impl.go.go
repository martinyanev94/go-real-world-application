package main

import "fmt"

// An interface that defines a contract
type Speaker interface {
    Speak() string
}

// A struct that implements the Speaker interface
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

// Another struct implementing the same interface
type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    var speaker Speaker

    speaker = Dog{}
    fmt.Println("Dog says:", speaker.Speak())

    speaker = Cat{}
    fmt.Println("Cat says:", speaker.Speak())
}
