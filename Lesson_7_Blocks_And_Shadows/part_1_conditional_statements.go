package main

import "fmt"

func main() {
    score := 85

    if score > 80 {
        fmt.Println("You passed with distinction!")
    } else {
        fmt.Println("You did not pass.")
    }
}
package main

import "fmt"

func main() {
    x := 5
    fmt.Println("Value of x before shadowing:", x)

    {
        x := 10 // This x shadows the outer x
        fmt.Println("Value of x inside the inner block:", x)
    }

    fmt.Println("Value of x after inner block:", x)
}
Value of x inside the inner block: 10
Value of x after inner block: 5
