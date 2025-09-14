package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	// Appending data to the buffer
	b.WriteString("Hello, ")
	b.WriteString("world!")
	
	// Converting the buffer to a string
	fmt.Println(b.String()) // Output: Hello, world!

	// Reusing the buffer by resetting it
	b.Reset()
	b.WriteString("Goodbye, world!")
	fmt.Println(b.String()) // Output: Goodbye, world!
}
