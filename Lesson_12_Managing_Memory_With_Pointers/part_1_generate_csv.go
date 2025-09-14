package main

import (
	"bytes"
	"fmt"
)

func createCSV(rows [][]string) string {
	var buffer bytes.Buffer

	for _, row := range rows {
		buffer.WriteString(fmt.Sprintf("%s\n", row[0])) // Simplistic CSV generation
	}

	return buffer.String()
}

func main() {
	data := [][]string{
		{"Header1", "Header2"},
		{"Row1Col1", "Row1Col2"},
		{"Row2Col1", "Row2Col2"},
	}

	csv := createCSV(data)
	fmt.Println(csv)
}
