slice := make([]int, 0, 1000) // Preallocating space for 1000 elements.
for i := 0; i < 1000; i++ {
    slice = append(slice, i)
}
fmt.Println(len(slice), cap(slice)) // Output: 1000 1000
