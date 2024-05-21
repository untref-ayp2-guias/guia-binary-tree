package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

// ejercicio2---------------------------------------
func SumaHojas(bt *binarytree.BinaryTree[int]) int {
	if bt.IsEmpty() {
		return 0
	}

	return sumaHojas(bt.GetRoot())
}

func sumaHojas(n *binarytree.BinaryNode[int]) int {
	return 0
}
