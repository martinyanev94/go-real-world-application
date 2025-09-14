type BinaryTree[T any] struct {
    Root *TreeNode[T]
}

func (bt *BinaryTree[T]) Insert(value T) {
    bt.Root = insertNode(bt.Root, value)
}

func insertNode[T any](node *TreeNode[T], value T) *TreeNode[T] {
    if node == nil {
        return &TreeNode[T]{Value: value}
    }
    if value < node.Value {
        node.Left = insertNode(node.Left, value)
    } else {
        node.Right = insertNode(node.Right, value)
    }
    return node
}
func (bt *BinaryTree[T]) InOrderTraverse(node *TreeNode[T]) {
    if node != nil {
        bt.InOrderTraverse(node.Left)
        fmt.Println(node.Value)
        bt.InOrderTraverse(node.Right)
    }
}
