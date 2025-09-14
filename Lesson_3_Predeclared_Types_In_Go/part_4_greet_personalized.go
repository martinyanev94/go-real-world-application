package main

import "fmt"

func main() {
    var greeting string = "Hello, World!"
    fmt.Println(greeting)

    // Concatenation of strings
    var name string = "Alice"
    var personalizedGreeting string = greeting + " It's great to see you, " + name + "."
    fmt.Println(personalizedGreeting)
}
