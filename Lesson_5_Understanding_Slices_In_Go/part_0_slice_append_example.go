slice := make([]int, 0, 5)
slice = append(slice, 10)
slice = append(slice, 20, 30)
slice = append(slice, 40)
fmt.Println(slice) // Output: [10 20 30 40]
