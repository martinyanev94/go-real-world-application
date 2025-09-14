// add.go
package main

func Add(x, y int) int {
    return x + y
}

// add_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(1, 2)
    expected := 3
    if result != expected {
        t.Errorf("Add(1, 2) = %d; want %d", result, expected)
    }
}

go
// add.go
package main

func Add(x, y int) int {
    return x + y
}

// add_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(1, 2)
    expected := 3
    if result != expected {
        t.Errorf("Add(1, 2) = %d; want %d", result, expected)
    }
}


go test

