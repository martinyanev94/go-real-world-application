intTree := BinaryTree[int]{}
intTree.Insert(5)
intTree.Insert(3)
intTree.Insert(7)
intTree.Insert(2)
intTree.Insert(4)

fmt.Println("In-order traversal of integer binary tree:")
intTree.InOrderTraverse(intTree.Root)
stringTree := BinaryTree[string]{}
stringTree.Insert("banana")
stringTree.Insert("apple")
stringTree.Insert("cherry")
stringTree.Insert("date")

fmt.Println("In-order traversal of string binary tree:")
stringTree.InOrderTraverse(stringTree.Root)
