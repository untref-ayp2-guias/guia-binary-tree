package binarytree

type BinaryTree[T any] struct {
	root *BinaryNode[T]
}

func NewBinaryTree[T any](data T, left *BinaryTree[T], right *BinaryTree[T]) *BinaryTree[T] {
	var node *BinaryNode[T]
	if left == nil && right == nil {
		node = NewBinaryNode(data, nil, nil)
	} else if right == nil {
		node = NewBinaryNode(data, left.root, nil)
	} else if left == nil {
		node = NewBinaryNode(data, nil, right.root)
	} else {
		node = NewBinaryNode(data, left.root, right.root)
	}
	return &BinaryTree[T]{root: node}
}

func (t *BinaryTree[T]) PrintPreOrder() {
	if t.root != nil {
		t.root.PrintPreOrder()
	}
}

func (t *BinaryTree[T]) PrintInOrder() {
	if t.root != nil {
		t.root.PrintInOrder()
	}
}

func (t *BinaryTree[T]) PrintPostOrder() {
	if t.root != nil {
		t.root.PrintPostOrder()
	}
}

func (t *BinaryTree[T]) Empty() {
	t.root = nil
}

func (t *BinaryTree[T]) IsEmpty() bool {
	return t.root == nil
}

func (t *BinaryTree[T]) Size() int {
	if t.root != nil {
		return t.root.Size()
	} else {
		return -1
	}
}

func (t *BinaryTree[T]) Height() int {
	if t.root != nil {
		return t.root.Height()
	} else {
		return 0
	}
}
