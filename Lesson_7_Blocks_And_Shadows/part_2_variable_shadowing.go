package main

import "fmt"

func main() {
    x := 15
    fmt.Println("Value of x before function call:", x)

    myFunction(x)

    fmt.Println("Value of x after function call:", x)
}

func myFunction(x int) {
    fmt.Println("Value of x inside function:", x)
    x = 30 // This modifies the shadowed variable, not the outer x
}
Value of x inside function: 15
Value of x after function call: 15
