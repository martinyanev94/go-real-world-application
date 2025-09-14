slice = append(slice[:1], slice[2:]...)
fmt.Println(slice) // Output will exclude the element at index 1
