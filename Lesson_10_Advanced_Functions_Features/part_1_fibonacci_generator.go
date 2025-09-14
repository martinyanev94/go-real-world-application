package main

import "fmt"

func fibonacciGenerator() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}

func main() {
    nextFibonacci := fibonacciGenerator()
    for i := 0; i < 10; i++ {
        fmt.Println(nextFibonacci())
    }
}
