package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

// ejercicio4---------------------------------------
func SumaPares(bt *binarytree.BinaryTree[int]) int {
	if bt.IsEmpty() {
		return 0
	}

	return sumaPares(bt.GetRoot())
}

func sumaPares(n *binarytree.BinaryNode[int]) int {
	return 0
}
