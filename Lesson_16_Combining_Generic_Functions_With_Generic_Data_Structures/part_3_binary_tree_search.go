func (bt *BinaryTree[T]) Search(value T) bool {
    return searchNode(bt.Root, value)
}

func searchNode[T comparable](node *TreeNode[T], value T) bool {
    if node == nil {
        return false
    }
    if value == node.Value {
        return true
    }
    if value < node.Value {
        return searchNode(node.Left, value)
    }
    return searchNode(node.Right, value)
}
