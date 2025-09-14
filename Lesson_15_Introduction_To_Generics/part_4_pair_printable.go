type Printable[T any] interface {
    String() string
}

type Pair[T Printable[T]] struct {
    val1, val2 T
}
