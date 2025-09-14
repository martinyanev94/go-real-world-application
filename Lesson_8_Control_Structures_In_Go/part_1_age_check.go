package main

import "fmt"

func main() {
    age := 18

    if age >= 18 {
        fmt.Println("You are an adult!")
    } else {
        fmt.Println("You are not an adult yet.")
    }
}
package main

import "fmt"

func main() {
    age := 16

    if age < 13 {
        fmt.Println("You are a child.")
    } else if age < 18 {
        fmt.Println("You are a teenager.")
    } else {
        fmt.Println("You are an adult!")
    }
}
