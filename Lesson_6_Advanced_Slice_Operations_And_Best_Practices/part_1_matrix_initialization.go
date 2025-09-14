rows, cols := 3, 4
matrix := make([][]int, rows) // Create a slice to hold slices
for i := range matrix {
    matrix[i] = make([]int, cols) // Populate each slice with a slice of length 'cols'
}

// Assign values to the matrix
matrix[0][0] = 1
matrix[1][1] = 2
matrix[2][2] = 3
for _, row := range matrix {
    fmt.Println(row) // Print each row of the matrix
}
