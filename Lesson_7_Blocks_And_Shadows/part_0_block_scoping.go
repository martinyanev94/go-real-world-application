package main

import "fmt"

func main() {
    fmt.Println("Outer block")

    // Declaring a variable in the outer block
    x := 10 
    fmt.Println("x from outer block:", x)

    {
        fmt.Println("Inner block")
        // Declaring a variable in the inner block
        y := 20 
        fmt.Println("y from inner block:", y)

        // Accessing the outer variable
        fmt.Println("Accessing x from inner block:", x)
    }

    // Trying to access y here will result in an error
    // fmt.Println("y from outer block:", y) // Uncommenting this line will cause a compilation error
}
