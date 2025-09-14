type Greetable interface {
    Greet() string
}
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s %s", p.FirstName, p.LastName)
}
func main() {
    var g Greetable = Person{FirstName: "Jane", LastName: "Doe"}
    fmt.Println(g.Greet())
}
