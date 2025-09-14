package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        fmt.Println("Outer loop iteration:", i)
        for j := 1; j <= 2; j++ {
            fmt.Println("  Inner loop iteration:", j)
        }
    }
}
package main

import "fmt"

func main() {
    count := 0
    for {
        count++
        fmt.Println("Count:", count)
        if count >= 5 {
            break
        }
    }
}
