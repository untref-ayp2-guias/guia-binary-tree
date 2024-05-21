package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

// ejercicio1---------------------------------------.

func SumaNodos(bt *binarytree.BinaryTree[int]) int {
	if bt.GetRoot() == nil {
		return 0
	}

	return sumaNodos(bt.GetRoot())
}

func sumaNodos(n *binarytree.BinaryNode[int]) int {
	return 0
}
