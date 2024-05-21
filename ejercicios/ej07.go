package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaHojasIzquierdas(t *binarytree.BinaryTree[int]) int {
	if t.IsEmpty() {
		return 0
	}

	return sumaHojasIzquierdas(t.GetRoot())
}

func sumaHojasIzquierdas(n *binarytree.BinaryNode[int]) int {
	return 0
}
