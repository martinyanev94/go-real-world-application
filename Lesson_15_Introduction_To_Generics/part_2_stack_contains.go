func (s *Stack[T]) Contains(val T) bool {
    for _, v := range s.vals {
        if v == val {
            return true
        }
    }
    return false
}
type Stack[T comparable] struct {
    vals []T
}
