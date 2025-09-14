func addNumbers(x, y int) (int, error) {
    if x < 0 || y < 0 {
        return 0, errors.New("negative input")
    }
    return x + y, nil
}

// Corresponding test
func Test_addNumbers_Error(t *testing.T) {
    _, err := addNumbers(-1, 3)
    if err == nil {
        t.Error("expected an error for negative input")
    }
}
