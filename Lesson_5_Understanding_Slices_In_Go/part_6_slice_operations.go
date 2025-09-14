anotherSlice := slice
anotherSlice[0] = 99
fmt.Println(slice[0]) // Output will also be 99
copySlice := make([]int, len(slice))
copy(copySlice, slice)
