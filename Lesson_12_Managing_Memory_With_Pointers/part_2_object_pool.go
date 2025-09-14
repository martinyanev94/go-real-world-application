package main

import (
	"fmt"
	"sync"
)

type Object struct {
	ID int
}

type ObjectPool struct {
	pool chan *Object
}

// NewObjectPool initializes a new memory pool
func NewObjectPool(size int) *ObjectPool {
	return &ObjectPool{
		pool: make(chan *Object, size),
	}
}

// Get retrieves an object from the pool or creates a new one
func (p *ObjectPool) Get() *Object {
	select {
	case obj := <-p.pool:
		return obj
	default:
		return &Object{}
	}
}

// Put returns an object to the pool
func (p *ObjectPool) Put(obj *Object) {
	select {
	case p.pool <- obj:
	default:
		// Pool is full, object is discarded
	}
}

func main() {
	pool := NewObjectPool(10)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			obj := pool.Get() // Get an object from the pool
			obj.ID = id
			fmt.Printf("Object ID: %d\n", obj.ID)
			pool.Put(obj) // Returning the object to the pool
		}(i)
	}
	wg.Wait()
}
