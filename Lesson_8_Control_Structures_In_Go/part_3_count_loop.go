package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Current count:", i)
    }
}
package main

import "fmt"

func main() {
    count := 1
    for count <= 5 {
        fmt.Println("Count is still:", count)
        count++
    }
}
