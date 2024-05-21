package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

// ejercicio3---------------------------------------
func SumaInternos(bt *binarytree.BinaryTree[int]) int {
	if bt.IsEmpty() {
		return 0
	}

	return sumaInternos(bt.GetRoot())
}

func sumaInternos(n *binarytree.BinaryNode[int]) int {
	return 0
}
