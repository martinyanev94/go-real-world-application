func processSlice(s []int) {
    for i := range s {
        s[i] *= 2 // Doubling each element
    }
}

numbers := []int{1, 2, 3, 4, 5}
processSlice(numbers)
fmt.Println(numbers) // Output: [2 4 6 8 10]
