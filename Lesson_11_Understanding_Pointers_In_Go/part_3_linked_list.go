package main

import "fmt"

type Node struct {
    value int
    next  *Node // Pointer to the next node
}

func main() {
    // Creating nodes
    node1 := &Node{value: 1}
    node2 := &Node{value: 2}
    node3 := &Node{value: 3}

    // Linking nodes
    node1.next = node2
    node2.next = node3

    // Traversing the linked list
    currentNode := node1
    for currentNode != nil {
        fmt.Println(currentNode.value)
        currentNode = currentNode.next // Move to the next node
    }
}
