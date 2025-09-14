package main

import "fmt"

// Public function
func PublicFunction() {
    fmt.Println("This is a public function.")
}

// Private function
func privateFunction() {
    fmt.Println("This is a private function.")
}

func main() {
    PublicFunction()   // Accessible
    privateFunction()  // Accessible within the same package
}
