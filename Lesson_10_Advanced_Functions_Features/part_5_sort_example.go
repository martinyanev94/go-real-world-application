package main

import "fmt"

func sort(arr []int, compare func(int, int) bool) []int {
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-1-i; j++ {
            if compare(arr[j+1], arr[j]) {
                arr[j], arr[j+1] = arr[j+1], arr[j] // Swap
            }
        }
    }
    return arr
}

func ascending(a, b int) bool {
    return a < b
}

func descending(a, b int) bool {
    return a > b
}

func main() {
    data := []int{5, 3, 8, 1, 2}
    fmt.Println("Ascending:", sort(data, ascending))
    fmt.Println("Descending:", sort(data, descending))
}
