func main() {
    p := Person{
        FirstName: "Jane",
        LastName:  "Doe",
        Age:       30,
    }
    fmt.Println("Name:", p.FirstName, p.LastName, "Age:", p.Age)
}
func (p Person) FullName() string {
    return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
func main() {
    p := Person{FirstName: "Jane", LastName: "Doe", Age: 30}
    fmt.Println("Full Name:", p.FullName())
}
