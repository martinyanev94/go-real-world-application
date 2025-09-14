package main

import "fmt"

func applyOperation(nums []int, operation func(int) int) []int {
    results := make([]int, len(nums))
    for i, num := range nums {
        results[i] = operation(num)
    }
    return results
}

func double(x int) int {
    return x * 2
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    doubled := applyOperation(numbers, double)
    fmt.Println("Doubled numbers:", doubled)
}
