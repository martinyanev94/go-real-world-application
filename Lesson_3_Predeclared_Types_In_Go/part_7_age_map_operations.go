package main

import "fmt"

func main() {
    // Create a map to store people's ages
    ages := make(map[string]int)

    // Adding key-value pairs to the map
    ages["Alice"] = 30
    ages["Bob"] = 25
    ages["Charlie"] = 35

    fmt.Println("Ages Map:", ages)

    // Accessing a value
    fmt.Printf("Alice's age is %d.\n", ages["Alice"])

    // Deleting a key from the map
    delete(ages, "Bob")
    fmt.Println("Updated Ages Map after deletion:", ages)
}
