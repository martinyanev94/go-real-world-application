original := []int{1, 2, 3}
copySlice := original
copySlice[0] = 99
fmt.Println(original) // Output: [99 2 3]
newSlice := make([]int, len(original))
copy(newSlice, original)
newSlice[0] = 100
fmt.Println(original) // Output: [1 2 3]
fmt.Println(newSlice) // Output: [100 2 3]
