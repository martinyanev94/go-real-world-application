fmt.Println(slice[1]) // Output: 20
slice[1] = 25
fmt.Println(slice) // Output: [10 25 30 40]
fmt.Println(slice[4]) // This will cause a runtime panic
