var i interface{}
i = "hello"
fmt.Println(i) // Works fine.
i = 42
fmt.Println(i) // Works fine too.
if str, ok := i.(string); ok {
    fmt.Println("String value:", str)
} else {
    fmt.Println("Not a string")
}
