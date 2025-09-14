type Person struct {
    Name string
    Age  int
}

people := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
}
fmt.Println(people[0].Name) // Output: Alice
people[1].Age = 26 // Update Bob's age
fmt.Println(people[1]) // Output: {Bob 26}
