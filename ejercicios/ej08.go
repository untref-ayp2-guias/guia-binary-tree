package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaDerechosImpares(bt *binarytree.BinaryTree[int]) int {
	if bt.IsEmpty() {
		return 0
	}

	return sumaDerechosImpares(bt.GetRoot())
}

func sumaDerechosImpares(n *binarytree.BinaryNode[int]) int {
	return 0
}
