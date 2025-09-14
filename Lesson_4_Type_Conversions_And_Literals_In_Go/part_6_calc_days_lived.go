package main

import (
    "fmt"
    "strconv"
)

func calcDaysLived(ageString string) int {
    age, err := strconv.Atoi(ageString) // Convert string to integer
    if err != nil {
        fmt.Println("Error converting age:", err)
        return 0
    }
    return age * 365 // Assuming every year has 365 days
}

func main() {
    ageInput := "30" // String literal for age
    daysLived := calcDaysLived(ageInput)
    fmt.Printf("You have lived for approximately %d days.\n", daysLived)
}
