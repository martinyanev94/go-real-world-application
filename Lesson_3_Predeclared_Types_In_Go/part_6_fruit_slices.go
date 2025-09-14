package main

import "fmt"

func main() {
    // Declare and initialize a slice
    fruits := []string{"Apple", "Banana", "Cherry"}

    // Append a new fruit to the slice
    fruits = append(fruits, "Dragonfruit")
    fmt.Println("Fruits Slice:", fruits)

    // Slicing to create a new slice
    tropicalFruits := fruits[1:3]
    fmt.Println("Tropical Fruits Slice:", tropicalFruits)
}
