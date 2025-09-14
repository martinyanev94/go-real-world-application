type Shape interface {
    Area() float64
}
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}
func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 6}
    
    PrintArea(c) // Outputs Area: 78.53981633974483
    PrintArea(r) // Outputs Area: 24
}
